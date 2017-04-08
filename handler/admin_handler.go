package handler

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type AdminHandler struct {
	Client client.Client
}

func (ah *AdminHandler) CorporationAllianceRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	//Upfront sanity check, don't do anything else until this is checked.
	//They could be anywhere... find them.  They SHOULD be at the same index of the marker enum
	//but first check that we have 2!
	if len(request.EntityType) != 2 || len(request.EntityId) != 2 || len(request.EntityName) != 2 || len(request.EntityTicker) != 2 {
		//TODO: Maybe I want some text to say WHY we failed?  Would an error object be more appropriate?
		response.Success = false
		return nil
	}

	alliance := findAllianceFromRequest(request)
	corporation := findCorpFromRequest(request)

	if alliance == nil || corporation == nil {
		response.Success = false
		return nil
	}

	//Now set the corporations alliance id
	corporation.AllianceId = &alliance.AllianceId

	//TODO: This is common
	role := repository.RoleRepo.FindByRoleName(request.Role)
	if role == nil {
		role = &model.Role{
			RoleName:         request.Role,
			ChatServiceGroup: request.Role,
		}

		err := repository.RoleRepo.Save(role)
		if err != nil {
			response.Success = false
			return nil
		}
	}

	//TODO: This is common
	dbAlliance := repository.AllianceRepo.FindByAllianceId(alliance.AllianceId)
	if dbAlliance == nil {
		err := repository.AllianceRepo.Save(alliance)
		if err != nil {
			response.Success = false
			return nil
		}
	}

	//TODO: This is common
	dbCorporation := repository.CorporationRepo.FindByCorporationId(corporation.CorporationId)
	if dbCorporation == nil {
		err := repository.CorporationRepo.Save(corporation)
		if err != nil {
			response.Success = false
			return nil
		}
	}

	err := repository.AccessRepo.SaveAllianceAndCorpRole(alliance.AllianceId, corporation.CorporationId, role)
	if err != nil {
		response.Success = false
		return nil
	}

	response.Success = true

	response.EntityType = make([]abaeve_auth.EntityType, 2)
	response.EntityType[0] = abaeve_auth.EntityType_ALLIANCE
	response.EntityType[1] = abaeve_auth.EntityType_CORPORATION

	response.EntityId = make([]int64, 2)
	response.EntityId[0] = alliance.AllianceId
	response.EntityId[1] = corporation.CorporationId

	response.EntityName = make([]string, 2)
	response.EntityName[0] = alliance.AllianceName
	response.EntityName[1] = corporation.CorporationName

	return nil
}

func (ah *AdminHandler) CorporationAllianceRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CharacterRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CharacterRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceCharacterLeadershipRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceCharacterLeadershipRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationCharacterLeadershipRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationCharacterLeadershipRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func findCorpFromRequest(request *abaeve_auth.AuthAdminRequest) *model.Corporation {
	var corporation *model.Corporation

	for idx, entityType := range request.EntityType {
		if entityType == abaeve_auth.EntityType_CORPORATION {
			corporation = &model.Corporation{
				CorporationId:     request.EntityId[idx],
				CorporationName:   request.EntityName[idx],
				CorporationTicker: request.EntityTicker[idx],
			}
		}
	}

	return corporation
}

func findAllianceFromRequest(request *abaeve_auth.AuthAdminRequest) *model.Alliance {
	var alliance *model.Alliance

	for idx, entityType := range request.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			alliance = &model.Alliance{
				AllianceId:     request.EntityId[idx],
				AllianceName:   request.EntityName[idx],
				AllianceTicker: request.EntityTicker[idx],
			}
		}
	}

	return alliance
}
