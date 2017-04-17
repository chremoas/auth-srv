package handler

import (
	"github.com/abaeve/auth-srv/proto"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityQueryHandler struct {
	Client client.Client
}

func (eqh *EntityQueryHandler) GetAlliances(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.Alliances) error {
	return nil
}

func (eqh *EntityQueryHandler) GetCorporations(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.Corporations) error {
	return nil
}

func (eqh *EntityQueryHandler) GetCharacters(context.Context, *abaeve_auth.EntityQueryRequest, *abaeve_auth.Characters) error {
	return nil
}
