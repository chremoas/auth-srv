package handler

import (
	"github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type EntityQueryHandler struct {
	Client client.Client
}

func (eqh *EntityQueryHandler) GetAlliances(ctx context.Context, request *abaeve_auth.EntityQueryRequest, response *abaeve_auth.AlliancesResponse) error {
	dbAliances := repository.AllianceRepo.FindAll()
	respAlliances := []*abaeve_auth.Alliance{}

	for _, alliance := range dbAliances {
		respAlliances = append(respAlliances,
			&abaeve_auth.Alliance{
				Id:     alliance.AllianceId,
				Name:   alliance.AllianceName,
				Ticker: alliance.AllianceTicker,
			},
		)
	}

	response.List = respAlliances

	return nil
}

func (eqh *EntityQueryHandler) GetCorporations(ctx context.Context, request *abaeve_auth.EntityQueryRequest, response *abaeve_auth.CorporationsResponse) error {
	dbCorporations := repository.CorporationRepo.FindAll()
	respCorporations := []*abaeve_auth.Corporation{}

	for _, corporation := range dbCorporations {
		respCorporation := &abaeve_auth.Corporation{
			Id:     corporation.CorporationId,
			Name:   corporation.CorporationName,
			Ticker: corporation.CorporationTicker,
		}

		if corporation.AllianceId != nil {
			respCorporation.AllianceId = *corporation.AllianceId
		}

		respCorporations = append(respCorporations, respCorporation)
	}

	response.List = respCorporations

	return nil
}

func (eqh *EntityQueryHandler) GetCharacters(ctx context.Context, request *abaeve_auth.EntityQueryRequest, response *abaeve_auth.CharactersResponse) error {
	dbCharacters := repository.CharacterRepo.FindAll()
	respCharacters := []*abaeve_auth.Character{}

	for _, character := range dbCharacters {
		respCharacters = append(respCharacters,
			&abaeve_auth.Character{
				Id:            character.CharacterId,
				Name:          character.CharacterName,
				CorporationId: character.CorporationId,
			},
		)
	}

	response.List = respCharacters

	return nil
}
