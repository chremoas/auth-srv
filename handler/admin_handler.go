package handler

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"github.com/abaeve/auth-srv/repository"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"errors"
)

type AdminHandler struct {
	Client client.Client
}

func (ah *AdminHandler) CorporationAllianceRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	//Upfront sanity check, don't do anything else until this is checked.
	//They could be anywhere... find them.  They SHOULD be at the same index of the marker enum
	//but first check that we have 2!
	//TODO: Maybe I want some text to say WHY we failed?  Would an error object be more appropriate?
	if !validateTwoEntitiesExist(request) {
		response.Success = false
		return nil
	}

	alliance, corporation, bothExist := findAllianceAndCorpFromRequest(request)
	if !bothExist {
		response.Success = false
		return nil
	}

	role, err := findOrSaveRoleByName(request.Role)
	if err != nil {
		response.Success = false
		return nil
	}

	err = findOrSaveAlliance(alliance)
	if err != nil {
		response.Success = false
		return nil
	}

	err = findOrSaveCorp(corporation)
	if err != nil {
		response.Success = false
		return nil
	}

	err = repository.AccessRepo.SaveAllianceAndCorpRole(alliance.AllianceId, corporation.CorporationId, role)
	if err != nil {
		response.Success = false
		return nil
	}

	//TODO: This is common and already starting to be copied and pasted
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
	if !validateTwoEntitiesExist(request) {
		response.Success = false
		return nil
	}

	alliance, corporation, bothExist := findAllianceAndCorpFromRequest(request)
	if !bothExist {
		response.Success = false
		return nil
	}

	role, err := findOrSaveRoleByName(request.Role)
	if err != nil {
		response.Success = false
		return nil
	}

	deletedRows, err := repository.AccessRepo.DeleteAllianceAndCorpRole(alliance.AllianceId, corporation.CorporationId, role)
	if err != nil {
		response.Success = false
		return nil
	}

	//TODO: I do need a more detailed error object inside the response
	if deletedRows > 1 {
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

func validateOneEntityExists(request *abaeve_auth.AuthAdminRequest) bool {
	if len(request.EntityType) != 1 || len(request.EntityId) != 1 || len(request.EntityName) != 1 || len(request.EntityTicker) != 1 {
		return false
	}

	return true
}

func validateTwoEntitiesExist(request *abaeve_auth.AuthAdminRequest) bool {
	if len(request.EntityType) != 2 || len(request.EntityId) != 2 || len(request.EntityName) != 2 || len(request.EntityTicker) != 2 {
		return false
	}

	return true
}

func findAllianceAndCorpFromRequest(request *abaeve_auth.AuthAdminRequest) (*model.Alliance, *model.Corporation, bool) {
	alliance := findAllianceFromRequest(request)
	corporation := findCorpFromRequest(request)

	if alliance == nil || corporation == nil {
		return nil, nil, false
	}

	//Now set the corporations alliance id
	corporation.AllianceId = &alliance.AllianceId

	return alliance, corporation, true
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

func findOrSaveRoleByName(roleName string) (*model.Role, error) {
	role := repository.RoleRepo.FindByRoleName(roleName)

	if role == nil {
		role = &model.Role{
			RoleName:         roleName,
			ChatServiceGroup: roleName,
		}

		err := repository.RoleRepo.Save(role)
		if err != nil {
			return nil, errors.New("Error while saving new role: " + err.Error())
		}
	}

	return role, nil
}

func findOrSaveAlliance(alliance *model.Alliance) error {
	dbAlliance := repository.AllianceRepo.FindByAllianceId(alliance.AllianceId)
	if dbAlliance == nil {
		err := repository.AllianceRepo.Save(alliance)
		if err != nil {
			return err
		}
	}

	return nil
}

func findOrSaveCorp(corporation *model.Corporation) error {
	dbCorporation := repository.CorporationRepo.FindByCorporationId(corporation.CorporationId)
	if dbCorporation == nil {
		err := repository.CorporationRepo.Save(corporation)
		if err != nil {
			return err
		}
	}

	return nil
}
