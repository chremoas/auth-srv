package handler

import (
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"golang.org/x/net/context"
	"testing"
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

	firstAllianceMatches := false
	secondAllianceMatches := false

	for _, alliance := range response.List {
		if alliance.Id == 1 {
			if alliance.Name == "Test Alliance 1" {
				if alliance.Ticker == "TSTA1" {
					firstAllianceMatches = true
				} else {
					t.Errorf("Expected (TSTA1) but received (%s)", alliance.Ticker)
				}
			} else {
				t.Errorf("Expected (Test Alliance 1) but received (%s)", alliance.Name)
			}
		} else if alliance.Id == 2 {
			if alliance.Name == "Test Alliance 2" {
				if alliance.Ticker == "TSTA2" {
					secondAllianceMatches = true
				} else {
					t.Errorf("Expected (TSTA2) but received (%s)", alliance.Ticker)
				}
			} else {
				t.Errorf("Expected (Test Alliance 2) but received (%s)", alliance.Name)
			}
		}
	}

	if !firstAllianceMatches {
		t.Error("First alliance had some issue's")
	}

	if !secondAllianceMatches {
		t.Error("Second alliance had some issue's")
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

	firstCorporationMatches := false
	secondCorporationMatches := false

	//Don't ask... I was in a weird mood...
	for _, corporation := range response.List {
		if corporation.Id == 1 {
			if corporation.Name == "Test Corporation 1" {
				if corporation.Ticker == "TSTC1" {
					if corporation.AllianceId == 1 {
						firstCorporationMatches = true
					} else {
						t.Errorf("Expected 1 but received (%d)", corporation.AllianceId)
					}
				} else {
					t.Errorf("Expected (TSTC1) but received (%s)", corporation.Ticker)
				}
			} else {
				t.Errorf("Expected (Test Corporation 1) but received (%s)", corporation.Name)
			}
		} else if corporation.Id == 2 {
			if corporation.Name == "Test Corporation 2" {
				if corporation.Ticker == "TSTC2" {
					if corporation.AllianceId == 2 {
						secondCorporationMatches = true
					} else {
						t.Errorf("Expected 2 but received (%d)", corporation.AllianceId)
					}
				} else {
					t.Errorf("Expected (TSTC2) but received (%s)", corporation.Ticker)
				}
			} else {
				t.Errorf("Expected (Test Corporation 2) but received (%s)", corporation.Name)
			}
		}
	}

	if !firstCorporationMatches {
		t.Error("First corporation had some issue's")
	}

	if !secondCorporationMatches {
		t.Error("Second corporation had some issue's")
	}
}

func TestEntityQueryHandler_GetCorporations_NoAlliance(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindAll().Return(
		[]*model.Corporation{
			{
				CorporationId:     int64(1),
				CorporationName:   "Test Corporation 1",
				CorporationTicker: "TSTC1",
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

	if len(response.List) != 1 {
		t.Errorf("Expected 1 but received (%d)", len(response.List))
	}

	if response.List[0].AllianceId != 0 {
		t.Errorf("Expected alliance id: (0) but received: (%d)", response.List[0].AllianceId)
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
		EntityType: abaeve_auth.EntityType_CHARACTER,
	}
	response := abaeve_auth.CharactersResponse{}

	err := entityQueryHandler.GetCharacters(context.Background(), &request, &response)

	if err != nil {
		t.Error("Received an error when none were expected")
	}

	if len(response.List) != 2 {
		t.Errorf("Expected 2 but received (%d)", len(response.List))
	}

	firstCharacterMatches := false
	secondCharacterMatches := false

	for _, character := range response.List {
		if character.Id == 1 {
			if character.Name == "Test Character 1" {
				if character.CorporationId == 1 {
					firstCharacterMatches = true
				} else {
					t.Errorf("Expected (1) but received (%d)", character.CorporationId)
				}
			} else {
				t.Errorf("Expected (Test Character 1) but received (%s)", character.Name)
			}
		} else if character.Id == 2 {
			if character.Name == "Test Character 2" {
				if character.CorporationId == 3 {
					secondCharacterMatches = true
				} else {
					t.Errorf("Expected (3) but received (%d)", character.CorporationId)
				}
			} else {
				t.Errorf("Expected (Test Character 2) but received (%s)", character.Name)
			}
		}
	}

	if !firstCharacterMatches {
		t.Error("First character had some issue's")
	}

	if !secondCharacterMatches {
		t.Error("Second character had some issue's")
	}
}
