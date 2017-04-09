package handler

import (
	"context"
	"errors"
	"github.com/abaeve/auth-srv/model"
	"github.com/abaeve/auth-srv/proto"
	"testing"
)

func TestAdminHandler_CorporationAllianceRoleAdd(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var allianceId int64
	allianceId = int64(1)

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	)
	mockAcceRepo.EXPECT().SaveAllianceAndCorpRole(
		int64(1),
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 2 {
		t.Fatalf("Expected 2 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithNewAlliance(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var allianceId int64
	allianceId = int64(1)

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(nil)
	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(nil)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	)
	mockAcceRepo.EXPECT().SaveAllianceAndCorpRole(
		int64(1),
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 2 {
		t.Fatalf("Expected 2 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithNewCorp(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var allianceId int64
	allianceId = int64(1)

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(nil)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	)
	mockAcceRepo.EXPECT().SaveAllianceAndCorpRole(
		int64(1),
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 2 {
		t.Fatalf("Expected 2 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithNewRole(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var allianceId int64
	allianceId = int64(1)

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(nil)
	mockRoleRepo.EXPECT().Save(
		&model.Role{
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).SetArg(
		0,
		model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	)
	mockAcceRepo.EXPECT().SaveAllianceAndCorpRole(
		int64(1),
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 2 {
		t.Fatalf("Expected 2 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithAllianceSaveError(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(nil)
	mockAlliRepo.EXPECT().Save(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	).Return(errors.New("Had a problem"))

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithCorpSaveError(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	var allianceId int64
	allianceId = int64(1)

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(nil)
	mockCorpRepo.EXPECT().Save(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
			AllianceId:        &allianceId,
		},
	).Return(errors.New("Had a prblem"))

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithRoleSaveError(t *testing.T) {
	mockCtrl, _, _, _, _, _, _, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(nil)
	mockRoleRepo.EXPECT().Save(
		&model.Role{
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).Return(errors.New("Had a problem"))

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithNotEnoughEntityStuff(t *testing.T) {
	mockCtrl, _, _, _, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE},
		EntityName:   []string{"Test Alliance"},
		EntityTicker: []string{"TSTA"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleAdd_WithTwoAlliances(t *testing.T) {
	mockCtrl, _, _, _, _, _, _, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_ALLIANCE},
		EntityName:   []string{"Test Alliance 1", "Test Alliance 2"},
		EntityTicker: []string{"TSTA1", "TSTA1"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleRemove(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAcceRepo.EXPECT().DeleteAllianceAndCorpRole(int64(1), int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).Return(1, nil)

	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 2 {
		t.Fatalf("Expected 2 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleRemove_WithTwoAlliances(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	mockAcceRepo.EXPECT().DeleteAllianceAndCorpRole(int64(1), int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).Return(1, nil).Times(0)

	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_ALLIANCE},
		EntityName:   []string{"Test Alliance 1", "Test Alliance 2"},
		EntityTicker: []string{"TSTA1", "TSTA2"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleRemove_With2Deletions(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAcceRepo.EXPECT().DeleteAllianceAndCorpRole(int64(1), int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).Return(2, nil)

	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_CorporationAllianceRoleRemove_WithError(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAcceRepo.EXPECT().DeleteAllianceAndCorpRole(int64(1), int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	).Return(0, errors.New("Had an issue!"))

	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1, 1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE, abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Alliance", "Test Corporation"},
		EntityTicker: []string{"TSTA", "TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationAllianceRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if response.Success {
		t.Fatal("Received a true success when false was expected")
	}

	if len(response.EntityId) != 0 {
		t.Fatalf("Expected 0 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if foundAllianceType {
		t.Fatal("Expected to find 0 response that was of type EntityType_ALLIANCE but found something")
	}

	if foundAllianceId != 0 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 0, foundAllianceId)
	}

	if foundCorpType {
		t.Fatal("Expected to find 0 response that was of type EntityType_CORPORATION but found something")
	}

	if foundCorpId != 0 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 0, foundCorpId)
	}
}

func TestAdminHandler_AllianceRoleAdd(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAlliRepo.EXPECT().FindByAllianceId(int64(1)).Return(
		&model.Alliance{
			AllianceId:     1,
			AllianceName:   "Test Alliance",
			AllianceTicker: "TSTA",
		},
	)
	mockAcceRepo.EXPECT().SaveAllianceRole(int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE},
		EntityName:   []string{"Test Alliance"},
		EntityTicker: []string{"TSTA"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.AllianceRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 1 {
		t.Fatalf("Expected 1 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}
}

func TestAdminHandler_AllianceRoleRemove(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAcceRepo.EXPECT().DeleteAllianceRole(int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_ALLIANCE},
		EntityName:   []string{"Test Alliance"},
		EntityTicker: []string{"TSTA"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.AllianceRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 1 {
		t.Fatalf("Expected 1 entity id's but received %d", len(response.EntityId))
	}

	foundAllianceType := false
	foundAllianceId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_ALLIANCE {
			foundAllianceType = true
			foundAllianceId = response.EntityId[idx]
		}
	}

	if !foundAllianceType {
		t.Fatal("Expected to find 1 response that was of type EntityType_ALLIANCE but found 0")
	} else if foundAllianceId != 1 {
		t.Fatalf("Expected alliance id: (%d) but received: (%d)", 1, foundAllianceId)
	}
}

func TestAdminHandler_CorporationRoleAdd(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockCorpRepo.EXPECT().FindByCorporationId(int64(1)).Return(
		&model.Corporation{
			CorporationId:     1,
			CorporationName:   "Test Corporation",
			CorporationTicker: "TSTC",
		},
	)
	mockAcceRepo.EXPECT().SaveCorporationRole(
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Corporation"},
		EntityTicker: []string{"TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationRoleAdd(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 1 {
		t.Fatalf("Expected 1 entity id's but received %d", len(response.EntityId))
	}
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CorporationRoleRemove(t *testing.T) {
	mockCtrl, _, _, _, _, _, mockAcceRepo, mockRoleRepo := SharedSetup(t)
	defer mockCtrl.Finish()

	mockRoleRepo.EXPECT().FindByRoleName("TEST_ROLE1").Return(
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)
	mockAcceRepo.EXPECT().DeleteCorporationRole(
		int64(1),
		&model.Role{
			RoleId:           1,
			RoleName:         "TEST_ROLE1",
			ChatServiceGroup: "TEST_ROLE1",
		},
	)

	adminHandler := &AdminHandler{}
	request := abaeve_auth.AuthAdminRequest{
		EntityId:     []int64{1},
		EntityType:   []abaeve_auth.EntityType{abaeve_auth.EntityType_CORPORATION},
		EntityName:   []string{"Test Corporation"},
		EntityTicker: []string{"TSTC"},
		Role:         "TEST_ROLE1",
	}
	response := abaeve_auth.AuthAdminResponse{}

	err := adminHandler.CorporationRoleRemove(context.Background(), &request, &response)

	if err != nil {
		t.Fatal("Received an error when one wasn't expected")
	}

	if !response.Success {
		t.Fatal("Received a false success when true was expected")
	}

	if len(response.EntityId) != 1 {
		t.Fatalf("Expected 1 entity id's but received %d", len(response.EntityId))
	}
	foundCorpType := false
	foundCorpId := int64(0)

	for idx, entityType := range response.EntityType {
		if entityType == abaeve_auth.EntityType_CORPORATION {
			foundCorpType = true
			foundCorpId = response.EntityId[idx]
		}
	}

	if !foundCorpType {
		t.Fatal("Expected to find 1 response that was of type EntityType_CORPORATION but found 0")
	} else if foundCorpId != 1 {
		t.Fatalf("Expected corp id: (%d) but received: (%d)", 1, foundCorpId)
	}
}

func TestAdminHandler_CharacterRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CharacterRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CharacterRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CharacterRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceCharacterLeadershipRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceCharacterLeadershipRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_AllianceCharacterLeadershipRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.AllianceCharacterLeadershipRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationCharacterLeadershipRoleAdd(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationCharacterLeadershipRoleAdd(context.Background(), &request, &response)

	t.Error("Implement me!")
}

func TestAdminHandler_CorporationCharacterLeadershipRoleRemove(t *testing.T) {
	adminHandler := &AdminHandler{}

	request := abaeve_auth.AuthAdminRequest{}
	response := abaeve_auth.AuthAdminResponse{}

	adminHandler.CorporationCharacterLeadershipRoleRemove(context.Background(), &request, &response)

	t.Error("Implement me!")
}
