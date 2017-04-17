package handler

import (
	"github.com/abaeve/auth-srv/proto"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityAdminHandler struct {
	Client client.Client
}

func (eah *EntityAdminHandler) AllianceUpdate(context.Context, *abaeve_auth.AllianceAdminRequest, *abaeve_auth.EntityAdminResponse) error {
	return nil
}

func (eah *EntityAdminHandler) CorporationUpdate(context.Context, *abaeve_auth.CoporationAdminRequest, *abaeve_auth.EntityAdminResponse) error {
	return nil
}

func (eah *EntityAdminHandler) CharacterUpdate(context.Context, *abaeve_auth.CharacterAdminRequest, *abaeve_auth.EntityAdminResponse) error {
	return nil
}
