package handler

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityAdminHandler struct {
	Client client.Client
}

func (eah *EntityAdminHandler) AllianceUpdate(ctx context.Context, request *abaeve_auth.AllianceAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Alliance == nil {
		response.Success = false
		response.ErrorText = "Invalid Alliance (nil)"
		return nil
	}

	if request.Alliance.Id == 0 {
		response.Success = false
		response.ErrorText = "Invalid Alliance Id (0/nil)"
		return nil
	}

	if len(request.Alliance.Name) == 0 {
		response.Success = false
		response.ErrorText = "Invalid Alliance Name (empty/nil)"
		return nil
	}

	if len(request.Alliance.Ticker) == 0 {
		response.Success = false
		response.ErrorText = "Invalid Alliance Ticker (empty/nil)"
		return nil
	}

	alliance := model.Alliance{
		AllianceId:     request.Alliance.Id,
		AllianceName:   request.Alliance.Name,
		AllianceTicker: request.Alliance.Ticker,
	}

	err := repository.AllianceRepo.Save(&alliance)

	if err != nil {
		response.Success = false
		response.ErrorText = "Error while saving: " + err.Error()
		return nil
	}

	response.Success = true
	return nil
}

func (eah *EntityAdminHandler) CorporationUpdate(ctx context.Context, request *abaeve_auth.CoporationAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Corporation == nil {
		response.Success = false
		response.ErrorText = "Invalid Corporation (nil)"
		return nil
	}

	if request.Corporation.Id == 0 {
		response.Success = false
		response.ErrorText = "Invalid Corporation Id (0/nil)"
		return nil
	}

	if len(request.Corporation.Name) == 0 {
		response.Success = false
		response.ErrorText = "Invalid Corporation Name (empty/nil)"
		return nil
	}

	if len(request.Corporation.Ticker) == 0 {
		response.Success = false
		response.ErrorText = "Invalid Corporation Ticker (empty/nil)"
		return nil
	}

	corporation := model.Corporation{
		CorporationId:     request.Corporation.Id,
		CorporationName:   request.Corporation.Name,
		CorporationTicker: request.Corporation.Ticker,
	}

	if request.Corporation.AllianceId != 0 {
		alliance := repository.AllianceRepo.FindByAllianceId(request.Corporation.AllianceId)
		if alliance == nil {
			response.Success = false
			response.ErrorText = "Invalid Alliance Id, Alliance doesn't exist"
			return nil
		} else {
			corporation.AllianceId = &request.Corporation.AllianceId
		}
	}

	err := repository.CorporationRepo.Save(&corporation)

	if err != nil {
		response.Success = false
		response.ErrorText = "Error while saving: " + err.Error()
		return nil
	}

	response.Success = true
	return nil
}

func (eah *EntityAdminHandler) CharacterUpdate(ctx context.Context, request *abaeve_auth.CharacterAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Character == nil {
		response.Success = false
		response.ErrorText = "Invalid Character (nil)"
		return nil
	}

	if request.Character.Id == 0 {
		response.Success = false
		response.ErrorText = "Invalid Character Id (0/nil)"
		return nil
	}

	if len(request.Character.Name) == 0 {
		response.Success = false
		response.ErrorText = "Invalid Character Name (empty/nil)"
		return nil
	}

	if request.Character.CorporationId == 0 {
		response.Success = false
		response.ErrorText = "Invalid Corporation Id (0/nil)"
		return nil
	}

	corporation := repository.CorporationRepo.FindByCorporationId(request.Character.CorporationId)
	if corporation == nil {
		response.Success = false
		response.ErrorText = "Invalid Corporation Id, Corporation doesn't exist"
		return nil
	}

	character := model.Character{
		CharacterId:   request.Character.Id,
		CharacterName: request.Character.Name,
		CorporationId: request.Character.CorporationId,
	}

	err := repository.CharacterRepo.Save(&character)

	if err != nil {
		response.Success = false
		response.ErrorText = "Error while saving: " + err.Error()
		return nil
	}

	response.Success = true
	return nil
}
