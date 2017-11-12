package mocks

//go:generate mockgen -package=mocks -destination=proto_mocks.go -imports=proto=github.com/abaeve/auth-srv/proto github.com/abaeve/auth-srv/proto UserAuthenticationClient,EntityQueryClient,UserAuthenticationAdminClient,EntityAdminClient
