package handler

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/micro/go-micro/client"
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
}

func (ah *AuthHandler) Create(context context.Context, request *abaeve_auth.AuthCreateRequest, response *abaeve_auth.AuthCreateResponse) error {
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
