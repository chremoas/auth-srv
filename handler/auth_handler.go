package handler

import (
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/abaeve/auth-srv/model"
	"crypto/rand"
	"encoding/hex"
)

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
			corporation.AllianceId = request.Alliance.Id
			corporation.Alliance = *alliance
		}

		err := repository.CorporationRepo.Save(corporation)

		if err != nil {
			return err
		}
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
	return nil
}
