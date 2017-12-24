package abaeve_auth

//go:generate pegomock generate --output=mocks/mock_userauthenticationclient.go --generate-matchers --package=authsrv_mocks UserAuthenticationClient
//go:generate pegomock generate --output=mocks/mock_entityqueryclient.go --generate-matchers --package=authsrv_mocks EntityQueryClient
//go:generate pegomock generate --output=mocks/mock_userauthenticationadminclient.go --generate-matchers --package=authsrv_mocks UserAuthenticationAdminClient
//go:generate pegomock generate --output=mocks/mock_entityadminclient.go --generate-matchers --package=authsrv_mocks EntityAdminClient
