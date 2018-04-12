package handler

import (
	"fmt"
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityAdminHandler struct {
	Client client.Client
}

//TODO: Do I really give a shit on delete whether or not the attributes besides id's are valid?
func (eah *EntityAdminHandler) AllianceUpdate(ctx context.Context, request *abaeve_auth.AllianceAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Alliance == nil {
		return fmt.Errorf("Invalid Alliance (nil)")
	}

	if request.Alliance.Id == 0 {
		return fmt.Errorf("Invalid Alliance Id (0/nil)")
	}

	if len(request.Alliance.Name) == 0 {
		return fmt.Errorf("Invalid Alliance Name (empty/nil)")
	}

	if len(request.Alliance.Ticker) == 0 {
		return fmt.Errorf("Invalid Alliance Ticker (empty/nil)")
	}

	if request.Operation == abaeve_auth.EntityOperation_ADD_OR_UPDATE {
		alliance := model.Alliance{
			AllianceId:     request.Alliance.Id,
			AllianceName:   request.Alliance.Name,
			AllianceTicker: request.Alliance.Ticker,
		}

		err := repository.AllianceRepo.Save(&alliance)
		if err != nil {
			//TODO: Find the consumers that may be expecting a non-error response and using the responses Success property and change that
			// Should we append to this error or send the raw error back?
			//response.ErrorText = "Error while saving: " + err.Error()
			return err
		}
	} else if request.Operation == abaeve_auth.EntityOperation_REMOVE {
		err := repository.AllianceRepo.Delete(request.Alliance.Id)
		if err != nil {
			return err
		}
	}

	response.Success = true
	return nil
}

func (eah *EntityAdminHandler) CorporationUpdate(ctx context.Context, request *abaeve_auth.CorporationAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Corporation == nil {
		return fmt.Errorf("Invalid Corporation (nil)")
	}

	if request.Corporation.Id == 0 {
		return fmt.Errorf("Invalid Corporation Id (0/nil)")
	}

	if len(request.Corporation.Name) == 0 {
		return fmt.Errorf("Invalid Corporation Name (empty/nil)")
	}

	if len(request.Corporation.Ticker) == 0 {
		return fmt.Errorf("Invalid Corporation Ticker (empty/nil)")
	}

	if request.Operation == abaeve_auth.EntityOperation_ADD_OR_UPDATE {
		corporation := model.Corporation{
			CorporationId:     request.Corporation.Id,
			CorporationName:   request.Corporation.Name,
			CorporationTicker: request.Corporation.Ticker,
		}

		if request.Corporation.AllianceId != 0 {
			alliance := repository.AllianceRepo.FindByAllianceId(request.Corporation.AllianceId)
			if alliance == nil {
				return fmt.Errorf("Invalid Alliance Id, Alliance doesn't exist")
			} else {
				corporation.AllianceId = &request.Corporation.AllianceId
			}
		}

		err := repository.CorporationRepo.Save(&corporation)
		if err != nil {
			//TODO: Find the consumers that may be expecting a non-error response and using the responses Success property and change that
			// Should we append to this error or send the raw error back?
			//response.ErrorText = "Error while saving: " + err.Error()
			return err
		}
	} else if request.Operation == abaeve_auth.EntityOperation_REMOVE {
		err := repository.CorporationRepo.Delete(request.Corporation.Id)
		if err != nil {
			return err
		}
	}

	response.Success = true
	return nil
}

func (eah *EntityAdminHandler) CharacterUpdate(ctx context.Context, request *abaeve_auth.CharacterAdminRequest, response *abaeve_auth.EntityAdminResponse) error {
	if request.Character == nil {
		return fmt.Errorf("Invalid Character (nil)")
	}

	if request.Character.Id == 0 {
		return fmt.Errorf("Invalid Character Id (0/nil)")
	}

	if len(request.Character.Name) == 0 {
		return fmt.Errorf("Invalid Character Name (empty/nil)")
	}

	if request.Character.CorporationId == 0 {
		return fmt.Errorf("Invalid Corporation Id (0/nil)")
	}

	if request.Operation == abaeve_auth.EntityOperation_ADD_OR_UPDATE {
		corporation := repository.CorporationRepo.FindByCorporationId(request.Character.CorporationId)
		if corporation == nil {
			return fmt.Errorf("Invalid Corporation Id, Corporation doesn't exist")
		}

		character := model.Character{
			CharacterId:   request.Character.Id,
			CharacterName: request.Character.Name,
			CorporationId: request.Character.CorporationId,
		}

		err := repository.CharacterRepo.Save(&character)

		if err != nil {
			//TODO: Find the consumers that may be expecting a non-error response and using the responses Success property and change that
			// Should we append to this error or send the raw error back?
			//response.ErrorText = "Error while saving: " + err.Error()
			return err
		}
	} else if request.Operation == abaeve_auth.EntityOperation_REMOVE {
		err := repository.CharacterRepo.Delete(request.Character.Id)
		if err != nil {
			return err
		}
	}

	response.Success = true
	return nil
}
