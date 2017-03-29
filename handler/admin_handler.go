package handler

import (
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/abaeve/auth-srv/proto"
)

type AdminHandler struct {
	Client client.Client
}

func (ah *AdminHandler) CharacterRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CharacterRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationAllianceRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationAllianceRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) CorporationRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceRoleAdd(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
	return nil
}

func (ah *AdminHandler) AllianceRoleRemove(context context.Context, request *abaeve_auth.AuthAdminRequest, response *abaeve_auth.AuthAdminResponse) error {
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
