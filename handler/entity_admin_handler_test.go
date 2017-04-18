package handler

import (
	"errors"
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"golang.org/x/net/context"
	"testing"
)

var iamSorryDave string = "I'm sorry, Dave. I'm afraid I can't do that."

func TestEntityAdminHandler_AllianceUpdate(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Return(nil)

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Id:     int64(1),
			Name:   "Test Alliance",
			Ticker: "TSTA",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if !response.Success {
		t.Error("Received unsuccessful when successful was expected")
	}
}

func TestEntityAdminHandler_AllianceUpdate_WithInvalidAlliance(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Alliance (nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_AllianceUpdate_WithInvalidAllianceId(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Name:   "Test Alliance",
			Ticker: "TSTA",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Alliance Id (0/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_AllianceUpdate_WithInvalidAllianceName(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Id:     int64(1),
			Ticker: "TSTA",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Alliance Name (empty/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_AllianceUpdate_WithInvalidAllianceTicker(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Id:   int64(1),
			Name: "Test Alliance",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Alliance Ticker (empty/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_AllianceUpdate_WithSaveError(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Return(errors.New(iamSorryDave))

	eah := EntityAdminHandler{}
	request := abaeve_auth.AllianceAdminRequest{
		Alliance: &abaeve_auth.Alliance{
			Id:     int64(1),
			Name:   "Test Alliance",
			Ticker: "TSTA",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.AllianceUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Error while saving: " + iamSorryDave

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Return(nil)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CoporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if !response.Success {
		t.Error("Received unsuccessful when successful was expected")
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithSaveError(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Return(errors.New(iamSorryDave))

	eah := EntityAdminHandler{}
	request := abaeve_auth.CoporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}
}

func TestEntityAdminHandler_CharacterUpdate(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(1),
		},
	).Return(nil)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			Name:          "Test Corporation",
			CorporationId: int64(1),
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if !response.Success {
		t.Error("Received unsuccessful when successful was expected")
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithSaveError(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCharRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     int64(1),
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Return(errors.New(iamSorryDave))

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:   int64(1),
			Name: "Test Corporation",
		},
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}
}
