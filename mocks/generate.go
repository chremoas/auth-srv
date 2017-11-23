package mocks

//go:generate mockgen -package=mocks -destination=proto_mocks.go -imports=proto=github.com/chremoas/auth-srv/proto github.com/chremoas/auth-srv/proto UserAuthenticationClient,EntityQueryClient,UserAuthenticationAdminClient,EntityAdminClient
