package handler

import (
	"errors"
	"github.com/chremoas/auth-srv/model"
	"github.com/chremoas/auth-srv/proto"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
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
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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

func TestEntityAdminHandler_AllianceUpdate_Delete(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	gomock.InOrder(
		mockAlliRepo.EXPECT().Delete(int64(1)).Return(nil),
		mockAlliRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while deleting the alliance")),
		mockAlliRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while updating corporations")),
	)

	eah := EntityAdminHandler{}

	Convey("Delete no error", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
			Alliance: &abaeve_auth.Alliance{
				Id: 1,
				Name:   "Test Alliance",
				Ticker: "TSTA",
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldBeNil)
		So(response, ShouldResemble, abaeve_auth.EntityAdminResponse{Success: true})
	})

	Convey("Delete with error deleting alliance", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
			Alliance: &abaeve_auth.Alliance{
				Id: 1,
				Name:   "Test Alliance",
				Ticker: "TSTA",
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while deleting the alliance")
	})

	Convey("Delete with error updating corporations", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.AllianceUpdate(context.Background(), &abaeve_auth.AllianceAdminRequest{
			Alliance: &abaeve_auth.Alliance{
				Id: 1,
				Name:   "Test Alliance",
				Ticker: "TSTA",
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while updating corporations")
	})
}

func TestEntityAdminHandler_CorporationUpdate(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Return(nil)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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

func TestEntityAdminHandler_CorporationUpdate_WithNoAlliance(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Times(0)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Return(nil)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:     int64(1),
			Name:   "Test Corporation",
			Ticker: "TSTC",
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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

func TestEntityAdminHandler_CorporationUpdate_WithInvalidCorp(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation (nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithInvalidCorpId(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Id (0/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithInvalidCorpName(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Name (empty/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithInvalidCorpTicker(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Times(0)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:   int64(1),
			CorporationName: "Test Corporation",
			AllianceId:      &allianceId,
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Ticker (empty/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithNonExistingAlliance(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(nil)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Alliance Id, Alliance doesn't exist"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CorporationUpdate_WithSaveError(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	allianceId := int64(1)

	mockAlliRepo.EXPECT().FindByAllianceId(allianceId).Return(
		&model.Alliance{
			AllianceId:     allianceId,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Return(errors.New(iamSorryDave))

	eah := EntityAdminHandler{}
	request := abaeve_auth.CorporationAdminRequest{
		Corporation: &abaeve_auth.Corporation{
			Id:         int64(1),
			Name:       "Test Corporation",
			Ticker:     "TSTC",
			AllianceId: allianceId,
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CorporationUpdate(context.Background(), &request, &response)

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

func TestEntityAdminHandler_CorporationUpdate_Delete(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	gomock.InOrder(
		mockCorpRepo.EXPECT().Delete(int64(1)).Return(nil),
		mockCorpRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while deleting the corporation")),
	)

	eah := EntityAdminHandler{}

	Convey("Delete no error", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
			Corporation: &abaeve_auth.Corporation{
				Id: 1,
				Name:   "Test Corp",
				Ticker: "TSTC",
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldBeNil)
		So(response, ShouldResemble, abaeve_auth.EntityAdminResponse{Success: true})
	})

	Convey("Delete with error deleting corporation", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CorporationUpdate(context.Background(), &abaeve_auth.CorporationAdminRequest{
			Corporation: &abaeve_auth.Corporation{
				Id: 1,
				Name:   "Test Corp",
				Ticker: "TSTC",
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while deleting the corporation")
	})
}

func TestEntityAdminHandler_CharacterUpdate(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	)
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
			Name:          "Test Character",
			CorporationId: int64(1),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
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

func TestEntityAdminHandler_CharacterUpdate_WithNoChar(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Times(0)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(1),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Character (nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithInvalidCharId(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Times(0)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(1),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Name:          "Test Character",
			CorporationId: int64(1),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Character Id (0/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithInvalidCharName(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Times(0)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(1),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			CorporationId: int64(1),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Character Name (empty/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithInvalidCorpId(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Times(0)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(0),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			Name:          "Test Character",
			CorporationId: int64(0),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Id (0/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithCorpIdNotSet(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	).Times(0)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(0),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:   int64(1),
			Name: "Test Character",
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Id (0/nil)"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithNonExistingCorp(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(nil)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(0),
		},
	).Times(0)

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			Name:          "Test Character",
			CorporationId: int64(1),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

	if err != nil {
		t.Errorf("Expected nil error but received: (%s)", err)
	}

	if response.Success {
		t.Error("Received successful when unsuccessful was expected")
	}

	expectedErrorText := "Invalid Corporation Id, Corporation doesn't exist"

	if response.ErrorText != expectedErrorText {
		t.Errorf("Expected: (%s) but received: (%s)", expectedErrorText, response.ErrorText)
	}
}

func TestEntityAdminHandler_CharacterUpdate_WithSaveError(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     int64(1),
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	)
	mockCharRepo.EXPECT().Save(
		&model.Character{
			CharacterId:   int64(1),
			CharacterName: "Test Character",
			CorporationId: int64(1),
		},
	).Return(errors.New(iamSorryDave))

	eah := EntityAdminHandler{}
	request := abaeve_auth.CharacterAdminRequest{
		Character: &abaeve_auth.Character{
			Id:            int64(1),
			Name:          "Test Character",
			CorporationId: int64(1),
		},
		Operation: abaeve_auth.EntityOperation_ADD_OR_UPDATE,
	}
	response := abaeve_auth.EntityAdminResponse{}

	err := eah.CharacterUpdate(context.Background(), &request, &response)

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

func TestEntityAdminHandler_CharacterUpdate_Delete(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	gomock.InOrder(
		mockCharRepo.EXPECT().Delete(int64(1)).Return(nil),
		mockCharRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while deleting the authentication codes")),
		mockCharRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while deleting the user_character_map")),
		mockCharRepo.EXPECT().Delete(int64(1)).Return(errors.New("dave I failed you while deleting the character")),
	)

	eah := EntityAdminHandler{}

	Convey("Delete no error", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
			Character: &abaeve_auth.Character{
				Id: 1,
				Name:   "Test Character",
				CorporationId: 1,
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldBeNil)
		So(response, ShouldResemble, abaeve_auth.EntityAdminResponse{Success: true})
	})

	Convey("Delete with error deleting authentication codes", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
			Character: &abaeve_auth.Character{
				Id: 1,
				Name:   "Test Character",
				CorporationId: 1,
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while deleting the authentication codes")
	})

	//The below two tests are actually redundant from the context of the AdminHandler but I'm putting them here for documentation purposes.
	Convey("Delete with error user_character_map ", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
			Character: &abaeve_auth.Character{
				Id: 1,
				Name:   "Test Character",
				CorporationId: 1,
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while deleting the user_character_map")
	})

	Convey("Delete with error deleting character", t, func() {
		response := abaeve_auth.EntityAdminResponse{}
		err := eah.CharacterUpdate(context.Background(), &abaeve_auth.CharacterAdminRequest{
			Character: &abaeve_auth.Character{
				Id: 1,
				Name:   "Test Character",
				CorporationId: 1,
			},
			Operation: abaeve_auth.EntityOperation_REMOVE,
		}, &response)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "dave I failed you while deleting the character")
	})
}
