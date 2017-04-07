package handler

import (
	"testing"
	"github.com/abaeve/auth-srv/proto"
	"context"
)

func TestAdminHandler_CorporationAllianceRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationAllianceRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationAllianceRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CharacterRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CharacterRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CharacterRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CharacterRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceCharacterLeadershipRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceCharacterLeadershipRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceCharacterLeadershipRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceCharacterLeadershipRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationCharacterLeadershipRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationCharacterLeadershipRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationCharacterLeadershipRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationCharacterLeadershipRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}
