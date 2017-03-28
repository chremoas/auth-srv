package handler

import (
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/abaeve/auth-srv/proto"
)

type AuthHandler struct {
	client client.Client
}

func (ah *AuthHandler) Create(context context.Context, request *abaeve_auth.AuthCreateRequest, response *abaeve_auth.AuthCreateResponse) error {
	return nil
}

func (ah *AuthHandler) Confirm(context context.Context, request *abaeve_auth.AuthConfirmRequest, response *abaeve_auth.AuthConfirmResponse) error {
	return nil
}
