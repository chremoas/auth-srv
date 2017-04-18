package handler

import (
	"testing"
	"github.com/abaeve/auth-srv/proto"
	"golang.org/x/net/context"
	"github.com/abaeve/auth-srv/model"
)

func TestEntityQueryHandler_GetAlliances(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().FindAll().Return(
		[]*model.Alliance{
			{
				AllianceId:     int64(1),
				AllianceName:   "Test Alliance 1",
				AllianceTicker: "TSTA1",
			},
			{
				AllianceId:     int64(2),
				AllianceName:   "Test Alliance 2",
				AllianceTicker: "TSTA2",
			},
		},
	)

	entityQueryHandler := EntityQueryHandler{}
	request := abaeve_auth.EntityQueryRequest{
		EntityType: abaeve_auth.EntityType_ALLIANCE,
	}
	response := abaeve_auth.AlliancesResponse{}

	err := entityQueryHandler.GetAlliances(context.Background(), &request, &response)

	if err != nil {
		t.Error("Received an error when none were expected")
	}

	if len(response.List) != 2 {
		t.Errorf("Expected 2 but received (%d)", len(response.List))
	}
}

func TestEntityQueryHandler_GetCorporations(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceIdOne := int64(1)
	allianceIdTwo := int64(2)

	mockCorpRepo.EXPECT().FindAll().Return(
		[]*model.Corporation{
			{
				CorporationId:     int64(1),
				CorporationName:   "Test Corporation 1",
				CorporationTicker: "TSTC1",
				AllianceId:        &allianceIdOne,
			},
			{
				CorporationId:     int64(2),
				CorporationName:   "Test Corporation 2",
				CorporationTicker: "TSTC2",
				AllianceId:        &allianceIdTwo,
			},
		},
	)

	entityQueryHandler := EntityQueryHandler{}
	request := abaeve_auth.EntityQueryRequest{
		EntityType: abaeve_auth.EntityType_CORPORATION,
	}
	response := abaeve_auth.CorporationsResponse{}

	err := entityQueryHandler.GetCorporations(context.Background(), &request, &response)

	if err != nil {
		t.Error("Received an error when none were expected")
	}

	if len(response.List) != 2 {
		t.Errorf("Expected 2 but received (%d)", len(response.List))
	}
}

func TestEntityQueryHandler_GetCharacters(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCharRepo.EXPECT().FindAll().Return(
		[]*model.Character{
			{
				CharacterId:   int64(1),
				CharacterName: "Test Character 1",
				CorporationId: int64(1),
			},
			{
				CharacterId:   int64(2),
				CharacterName: "Test Character 2",
				CorporationId: int64(3),
			},
		},
	)

	entityQueryHandler := EntityQueryHandler{}
	request := abaeve_auth.EntityQueryRequest{
		EntityType: abaeve_auth.EntityType_ALLIANCE,
	}
	response := abaeve_auth.AlliancesResponse{}

	err := entityQueryHandler.GetAlliances(context.Background(), &request, &response)

	if err != nil {
		t.Error("Received an error when none were expected")
	}

	if len(response.List) != 2 {
		t.Errorf("Expected 2 but received (%d)", len(response.List))
	}
}
