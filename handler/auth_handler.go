package handler

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	rolesrv "github.com/chremoas/role-srv/proto"
	"github.com/chremoas/services-common/config"
	"github.com/chremoas/services-common/sets"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type authError struct {
	message string
}

func (ae *authError) Error() string {
	return ae.message
}

type AuthHandler struct {
	Client client.Client
	Logger *zap.Logger
}

type clientList struct {
	roles rolesrv.RolesClient
}

var clients clientList

func NewAuthHandler(config *config.Configuration, service micro.Service, log *zap.Logger) abaeve_auth.UserAuthenticationHandler {
	c := service.Client()

	clients = clientList{
		roles: rolesrv.NewRolesClient(config.LookupService("srv", "role"), c),
	}

	return &AuthHandler{Client: c, Logger: log}
}

func (ah *AuthHandler) Create(context context.Context, request *abaeve_auth.AuthCreateRequest, response *abaeve_auth.AuthCreateResponse) error {
	ah.Logger.Info("Call to Create()")

	var alliance *model.Alliance

	//We MIGHT NOT have any kind of alliance information
	if request.Alliance != nil {
		alliance = repository.AllianceRepo.FindByAllianceId(request.Alliance.Id)

		if alliance == nil {
			alliance = &model.Alliance{
				AllianceId:     request.Alliance.Id,
				AllianceName:   request.Alliance.Name,
				AllianceTicker: request.Alliance.Ticker,
			}
			err := repository.AllianceRepo.Save(alliance)

			if err != nil {
				return err
			}

			publication := ah.Client.NewPublication(abaeve_auth.AllianceAddTopic(), request.Alliance)
			ah.Client.Publish(context, publication)
		}
	}

	corporation := repository.CorporationRepo.FindByCorporationId(request.Corporation.Id)

	if corporation == nil {
		corporation = &model.Corporation{
			CorporationId:     request.Corporation.Id,
			CorporationName:   request.Corporation.Name,
			CorporationTicker: request.Corporation.Ticker,
		}

		if alliance != nil {
			corporation.AllianceId = &request.Alliance.Id
			corporation.Alliance = *alliance
		}

		err := repository.CorporationRepo.Save(corporation)

		if err != nil {
			return err
		}

		publication := ah.Client.NewPublication(abaeve_auth.CorporationAddTopic(), request.Corporation)
		ah.Client.Publish(context, publication)
	}

	character := repository.CharacterRepo.FindByCharacterId(request.Character.Id)

	if character == nil {
		character = &model.Character{
			CharacterId:   request.Character.Id,
			CharacterName: request.Character.Name,
			Token:         request.Token,
			CorporationId: request.Corporation.Id,
			Corporation:   *corporation,
		}
		err := repository.CharacterRepo.Save(character)

		if err != nil {
			return err
		}

		publication := ah.Client.NewPublication(abaeve_auth.CharacterAddTopic(), request.Character)
		ah.Client.Publish(context, publication)
	}

	//Now... make an auth string... hopefully this isn't too ugly
	b := make([]byte, 6)
	rand.Read(b)
	authCode := hex.EncodeToString(b)

	err := repository.AuthenticationCodeRepo.Save(character, authCode)

	if err != nil {
		return err
	}

	response.AuthenticationCode = authCode
	response.Success = true

	return nil
}

func (ah *AuthHandler) Confirm(context context.Context, request *abaeve_auth.AuthConfirmRequest, response *abaeve_auth.AuthConfirmResponse) error {
	ah.Logger.Info("Call to Confirm()")

	character := repository.CharacterRepo.FindByAutenticationCode(request.AuthenticationCode)

	if character == nil {
		return &authError{message: "Invalid Auth Attempt"}
	}

	user := repository.UserRepo.FindByChatId(request.UserId)

	if user == nil {
		user = &model.User{ChatId: request.UserId}
		err := repository.UserRepo.Save(user)

		if err != nil {
			return errors.New("Error saving user: " + err.Error())
		}
	}

	err := repository.UserRepo.LinkCharacterToUserByAuthCode(request.AuthenticationCode, user)

	if err != nil {
		return errors.New("Error linking user: " + err.Error())
	}

	response.Success = true
	response.CharacterName = character.CharacterName

	return nil
}

func (ah *AuthHandler) SyncToRoleService(ctx context.Context, request *abaeve_auth.NilRequest, response *abaeve_auth.SyncToRoleResponse) error {
	var allianceMembers = make(map[string]*sets.StringSet)
	var corpMembers = make(map[string]*sets.StringSet)
	var allianceSet = sets.NewStringSet()
	var corpSet = sets.NewStringSet()
	var filterSet = sets.NewStringSet()

	sugar := ah.Logger.Sugar()
	sugar.Info("Call to SyncToRoleService()")

	filters, err := clients.roles.GetFilters(ctx, &rolesrv.NilMessage{})
	if err != nil {
		return err
	}

	for f := range filters.FilterList {
		filterSet.Add(filters.FilterList[f].Name)
	}

	authMembers, err := repository.AccessRepo.GetMembership()
	if err != nil {
		return err
	}

	// Check if the filters exist, if they don't, create them
	for m := range authMembers {
		if _, ok := allianceMembers[authMembers[m].AllianceTicker]; !ok {
			allianceMembers[authMembers[m].AllianceTicker] = sets.NewStringSet()
		}

		if _, ok := corpMembers[authMembers[m].CorpTicker]; !ok {
			corpMembers[authMembers[m].CorpTicker] = sets.NewStringSet()
		}

		// Why doesn't response get updated?
		if !filterSet.Contains(authMembers[m].AllianceTicker) {
			ah.addFilter(ctx, authMembers[m].AllianceTicker, authMembers[m].AllianceName, response)
		}

		if !filterSet.Contains(authMembers[m].CorpTicker) {
			ah.addFilter(ctx, authMembers[m].CorpTicker, authMembers[m].CorpName, response)
		}
	}

	for m := range authMembers {
		allianceMembers[authMembers[m].AllianceTicker].Add(authMembers[m].ChatId)
		corpMembers[authMembers[m].CorpTicker].Add(authMembers[m].ChatId)
	}

	for a := range allianceMembers {
		members, err := clients.roles.GetMembers(ctx, &rolesrv.Filter{Name: a})
		if err != nil {
			return err
		}

		allianceSet.FromSlice(members.Members)
		clients.roles.AddMembers(ctx, &rolesrv.Members{Filter: a, Name: allianceMembers[a].Difference(allianceSet).ToSlice()})
		clients.roles.RemoveMembers(ctx, &rolesrv.Members{Filter: a, Name: allianceSet.Difference(allianceMembers[a]).ToSlice()})
	}

	for c := range corpMembers {
		members, err := clients.roles.GetMembers(ctx, &rolesrv.Filter{Name: c})
		if err != nil {
			return err
		}

		corpSet.FromSlice(members.Members)
		clients.roles.AddMembers(ctx, &rolesrv.Members{Filter: c, Name: corpMembers[c].Difference(corpSet).ToSlice()})
		clients.roles.RemoveMembers(ctx, &rolesrv.Members{Filter: c, Name: corpSet.Difference(corpMembers[c]).ToSlice()})
	}

	clients.roles.SyncRoles(ctx, &rolesrv.NilMessage{})
	return nil
}

func (ah AuthHandler) addFilter(ctx context.Context, name string, description string, response *abaeve_auth.SyncToRoleResponse) {
	sugar := ah.Logger.Sugar()
	sugar.Infof("Adding filter '%s': %s", name, description)

	clients.roles.AddFilter(ctx, &rolesrv.Filter{
		Name:        name,
		Description: description,
	})

	response.Roles = append(response.Roles, &abaeve_auth.Role{
		Name:        name,
		Description: description,
	})
}
