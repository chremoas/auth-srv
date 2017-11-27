package mocks

//go:generate mockgen -package=mocks -destination=proto_mocks.go -imports=proto=github.com/chremoas/auth-srv/proto github.com/chremoas/auth-srv/proto UserAuthenticationClient,EntityQueryClient,UserAuthenticationAdminClient,EntityAdminClient
//go:generate mockgen -package=mocks -destination=repository_mocks.go github.com/chremoas/auth-srv/repository AllianceRepository,CorporationRepository,CharacterRepository,UserRepository,RoleRepository,AuthenticationCodeRepository
