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
	}

	response.Success = true
	return nil
}

func (eah *EntityAdminHandler) CorporationUpdate(ctx context.Context, request *abaeve_auth.CoporationAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	return nil
}

func (eah *EntityAdminHandler) CharacterUpdate(ctx context.Context, request *abaeve_auth.CharacterAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	return nil
}
