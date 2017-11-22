package handler

import (
	"errors"
	"github.com/chremoas/auth-srv/mocks"
	"github.com/chremoas/auth-srv/model"
	proto "github.com/chremoas/auth-srv/proto"
	"github.com/chremoas/auth-srv/repository"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"testing"
)

type mockPublication struct {
	topic       string
	message     interface{}
	contentType string
}

func (mp mockPublication) Topic() string {
	return mp.topic
}

func (mp mockPublication) Message() interface{} {
	return mp.message
}

func (mp mockPublication) ContentType() string {
	return mp.contentType
}

func SharedSetup(t *testing.T) (*gomock.Controller,
	*mocks.MockAuthenticationCodeRepository,
	*mocks.MockUserRepository,
	*mocks.MockCharacterRepository,
	*mocks.MockCorporationRepository,
	*mocks.MockAllianceRepository,
	*mocks.MockAccessesRepository,
	*mocks.MockRoleRepository) {
	mockCtrl := gomock.NewController(t)

	repository.AuthenticationCodeRepo = mocks.NewMockAuthenticationCodeRepository(mockCtrl)
	repository.UserRepo = mocks.NewMockUserRepository(mockCtrl)
	repository.CharacterRepo = mocks.NewMockCharacterRepository(mockCtrl)
	repository.CorporationRepo = mocks.NewMockCorporationRepository(mockCtrl)
	repository.AllianceRepo = mocks.NewMockAllianceRepository(mockCtrl)
	repository.AccessRepo = mocks.NewMockAccessesRepository(mockCtrl)
	repository.RoleRepo = mocks.NewMockRoleRepository(mockCtrl)

	mockAuthRepo := repository.AuthenticationCodeRepo.(*mocks.MockAuthenticationCodeRepository)
	mockUserRepo := repository.UserRepo.(*mocks.MockUserRepository)
	mockCharRepo := repository.CharacterRepo.(*mocks.MockCharacterRepository)
	mockCorpRepo := repository.CorporationRepo.(*mocks.MockCorporationRepository)
	mockAlliRepo := repository.AllianceRepo.(*mocks.MockAllianceRepository)
	mockAcceRepo := repository.AccessRepo.(*mocks.MockAccessesRepository)
	mockRoleRepo := repository.RoleRepo.(*mocks.MockRoleRepository)

	return mockCtrl, mockAuthRepo, mockUserRepo, mockCharRepo, mockCorpRepo, mockAlliRepo, mockAcceRepo, mockRoleRepo
}

func TestCreateEmptyDb(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockAlliRepo.EXPECT().Save(
			&model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.AllianceAddTopic(),
			authCreateRequest.Alliance,
		).Return(mockPublication{
			message:     authCreateRequest.Alliance,
			topic:       "AllianceAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Alliance,
			topic:       "AllianceAdd",
			contentType: "ContentType",
		}),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CorporationAddTopic(),
			authCreateRequest.Corporation,
		).Return(mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Return(mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestCreateNoAllianceCorporation(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockAlliRepo.EXPECT().Save(
			&model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.AllianceAddTopic(),
			authCreateRequest.Alliance,
		).Return(mockPublication{
			message:     authCreateRequest.Alliance,
			topic:       "AllianceAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Alliance,
			topic:       "AllianceAdd",
			contentType: "ContentType",
		}),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CorporationAddTopic(),
			authCreateRequest.Corporation,
		).Return(mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Return(mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceExistsNoCorpOrChar(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(0),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(nil),
		mockCorpRepo.EXPECT().Save(
			&model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CorporationAddTopic(),
			authCreateRequest.Corporation,
		).Return(mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}).Times(1),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Corporation,
			topic:       "CorporationAdd",
			contentType: "ContentType",
		}).Times(1),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Return(mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}).Times(1),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}).Times(1),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceAndCorpExistButNoChar(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(0),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(&model.Corporation{
			CorporationId:     authCreateRequest.Corporation.Id,
			CorporationName:   authCreateRequest.Corporation.Name,
			CorporationTicker: authCreateRequest.Corporation.Ticker,
			AllianceId:        &authCreateRequest.Alliance.Id,
			Alliance: model.Alliance{
				AllianceId:     authCreateRequest.Alliance.Id,
				AllianceName:   authCreateRequest.Alliance.Name,
				AllianceTicker: authCreateRequest.Alliance.Ticker,
			},
		}),
		mockCorpRepo.EXPECT().Save(&model.Corporation{}).Times(0),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(0),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(0),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(nil),
		mockCharRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
		).Return(nil),
		mockClient.EXPECT().NewPublication(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Return(mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}).Times(1),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Character,
			topic:       "CharacterAdd",
			contentType: "ContentType",
		}).Times(1),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceAndCorpAndCharExist(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(&model.Alliance{
			AllianceId:     authCreateRequest.Alliance.Id,
			AllianceName:   authCreateRequest.Alliance.Name,
			AllianceTicker: authCreateRequest.Alliance.Ticker,
		}),
		mockAlliRepo.EXPECT().Save(&model.Alliance{}).Times(0),
		mockClient.EXPECT().NewPublication(proto.AllianceAddTopic(),
			authCreateRequest.Alliance,
		).Times(0),
		mockClient.EXPECT().Publish(ctx, mockPublication{
			message:     authCreateRequest.Alliance,
			topic:       "AllianceAdd",
			contentType: "ContentType",
		}).Times(0),

		mockCorpRepo.EXPECT().FindByCorporationId(authCreateRequest.Corporation.Id).Return(&model.Corporation{
			CorporationId:     authCreateRequest.Corporation.Id,
			CorporationName:   authCreateRequest.Corporation.Name,
			CorporationTicker: authCreateRequest.Corporation.Ticker,
		}),
		mockCorpRepo.EXPECT().Save(&model.Corporation{}).Times(0),
		mockClient.EXPECT().NewPublication(proto.CorporationAddTopic(),
			authCreateRequest.Corporation,
		).Times(0),
		mockClient.EXPECT().Publish(proto.CorporationAddTopic(),
			authCreateRequest.Corporation,
		).Times(0),

		mockCharRepo.EXPECT().FindByCharacterId(authCreateRequest.Character.Id).Return(&model.Character{
			CharacterId:   authCreateRequest.Character.Id,
			CharacterName: authCreateRequest.Character.Name,
			Token:         authCreateRequest.Token,
			CorporationId: authCreateRequest.Corporation.Id,
			Corporation: model.Corporation{
				CorporationId:     authCreateRequest.Corporation.Id,
				CorporationName:   authCreateRequest.Corporation.Name,
				CorporationTicker: authCreateRequest.Corporation.Ticker,
				AllianceId:        &authCreateRequest.Alliance.Id,
				Alliance: model.Alliance{
					AllianceId:     authCreateRequest.Alliance.Id,
					AllianceName:   authCreateRequest.Alliance.Name,
					AllianceTicker: authCreateRequest.Alliance.Ticker,
				},
			},
		}),
		mockCharRepo.EXPECT().Save(&model.Character{}).Times(0),
		mockClient.EXPECT().NewPublication(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Times(0),
		mockClient.EXPECT().Publish(proto.CharacterAddTopic(),
			authCreateRequest.Character,
		).Times(0),

		mockAuthRepo.EXPECT().Save(
			&model.Character{
				CharacterId:   authCreateRequest.Character.Id,
				CharacterName: authCreateRequest.Character.Name,
				Token:         authCreateRequest.Token,
				CorporationId: authCreateRequest.Corporation.Id,
				Corporation: model.Corporation{
					CorporationId:     authCreateRequest.Corporation.Id,
					CorporationName:   authCreateRequest.Corporation.Name,
					CorporationTicker: authCreateRequest.Corporation.Ticker,
					AllianceId:        &authCreateRequest.Alliance.Id,
					Alliance: model.Alliance{
						AllianceId:     authCreateRequest.Alliance.Id,
						AllianceName:   authCreateRequest.Alliance.Name,
						AllianceTicker: authCreateRequest.Alliance.Ticker,
					},
				},
			},
			gomock.Any(),
		).Return(nil),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err != nil {
		t.Fatalf("Received an error on Create call, expected nothing: %s", err)
	}

	if authCreateResponse.AuthenticationCode == "" {
		t.Fatal("Expected at least something as an authentication code, got nothing")
	}
}

func TestAllianceErrorCondition(t *testing.T) {
	mockCtrl, _, _, _, _, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(authCreateRequest.Alliance.Id).Return(nil),
		mockClient.EXPECT().NewPublication(proto.AllianceAddTopic(),
			authCreateRequest.Alliance,
		).Times(0),
		mockClient.EXPECT().Publish(proto.AllianceAddTopic(),
			authCreateRequest.Alliance,
		).Times(0),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Return(errors.New("Don't do that alliance")),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that alliance" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestCorporationErrorCondition(t *testing.T) {
	mockCtrl, _, _, _, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Return(errors.New("Don't do that corp")),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that corp" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestCharacterErrorCondition(t *testing.T) {
	mockCtrl, _, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Times(1),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockCharRepo.EXPECT().FindByCharacterId(gomock.Any()).Return(nil),
		mockCharRepo.EXPECT().Save(gomock.Any()).Return(errors.New("Don't do that char")),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that char" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestAuthCodeErrorCondition(t *testing.T) {
	mockCtrl, mockAuthRepo, _, mockCharRepo, mockCorpRepo, mockAlliRepo, _, _ := SharedSetup(t)
	mockClient := mocks.NewMockClient(mockCtrl)
	defer mockCtrl.Finish()

	authCreateRequest := proto.AuthCreateRequest{
		Token:       "mytoken",
		Alliance:    &proto.Alliance{Name: "Test Alliance", Id: 1, Ticker: "TSTA"},
		Corporation: &proto.Corporation{Name: "Test Corp", Id: 1, Ticker: "TSTC"},
		Character:   &proto.Character{Name: "Test Character", Id: 1},
		AuthScope:   []string{"scope1", "scope2"},
	}
	var authCreateResponse proto.AuthCreateResponse
	var ctx context.Context

	authHandler := AuthHandler{mockClient}

	//Set our expectations
	gomock.InOrder(
		mockAlliRepo.EXPECT().FindByAllianceId(gomock.Any()).Return(nil),
		mockAlliRepo.EXPECT().Save(gomock.Any()).Times(1),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockCorpRepo.EXPECT().FindByCorporationId(gomock.Any()).Return(nil),
		mockCorpRepo.EXPECT().Save(gomock.Any()).Times(1),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockCharRepo.EXPECT().FindByCharacterId(gomock.Any()).Return(nil),
		mockCharRepo.EXPECT().Save(gomock.Any()).Return(nil),
		mockClient.EXPECT().NewPublication(gomock.Any(), gomock.Any()).Times(1),
		mockClient.EXPECT().Publish(gomock.Any(), gomock.Any()).Times(1),

		mockAuthRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errors.New("Don't do that auth")),
	)

	err := authHandler.Create(ctx, &authCreateRequest, &authCreateResponse)

	if err == nil && err.Error() == "Don't do that auth" {
		t.Fatal("Expected an error but got nothing")
	}
}

func TestConfirmWithNoChar(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(nil),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Times(0),
		mockUserRepo.EXPECT().Save(gomock.Any()).Times(0),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(gomock.Any(), gomock.Any()).Times(0),
		mockAcceRepo.EXPECT().FindByChatId(gomock.Any()).Times(0),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	if err == nil || err.Error() != "Invalid Auth Attempt" {
		t.Errorf("Expected a specific error but received: %s", err)
	}
}

func TestConfirmWithAuthNoUser(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo, _ := SharedSetup(t)

	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}
	expectedCharName := "Test Character Result"

	allianceId := int64(1)

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(
			&model.Character{
				CharacterId:   1,
				CharacterName: expectedCharName,
				CorporationId: 1,
				Corporation: model.Corporation{
					CorporationId:     1,
					CorporationName:   "Test Corporation Result",
					CorporationTicker: "TSTC",
					AllianceId:        &allianceId,
					Alliance: model.Alliance{
						AllianceId:     1,
						AllianceName:   "Test Alliance Result",
						AllianceTicker: "TSTA",
					},
				},
			},
		),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil),
		mockUserRepo.EXPECT().Save(
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(authConfirmRequest.AuthenticationCode,
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		).Return(nil),
		mockAcceRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return([]string{"ROLE1", "ROLE2", "ROLE3"}, nil),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	if err != nil {
		t.Errorf("Expected no error but received: %s", err)
	}

	if len(authConfirmResponse.CharacterName) == 0 || authConfirmResponse.CharacterName == "" {
		t.Errorf("Response character name: (%s) doesn't match expected: (%s)", authConfirmResponse.CharacterName, expectedCharName)
	}
}

func TestConfirm_WithUserSaveError(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo, _ := SharedSetup(t)

	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}
	expectedCharName := "Test Character Result"

	allianceId := int64(1)

	underlyingError := "I'm sorry, Dave. I'm afraid I can't do that."

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(
			&model.Character{
				CharacterId:   1,
				CharacterName: expectedCharName,
				CorporationId: 1,
				Corporation: model.Corporation{
					CorporationId:     1,
					CorporationName:   "Test Corporation Result",
					CorporationTicker: "TSTC",
					AllianceId:        &allianceId,
					Alliance: model.Alliance{
						AllianceId:     1,
						AllianceName:   "Test Alliance Result",
						AllianceTicker: "TSTA",
					},
				},
			},
		),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil),
		mockUserRepo.EXPECT().Save(
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		).Return(errors.New(underlyingError)),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(authConfirmRequest.AuthenticationCode,
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		).Return(nil).Times(0),
		mockAcceRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return([]string{"ROLE1", "ROLE2", "ROLE3"}, nil).Times(0),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	expectedError := "Error saving user: " + underlyingError

	if err != nil && err.Error() != expectedError {
		t.Errorf("Expected error message (%s) but received (%s)", expectedError, err)
	} else if err == nil {
		t.Error("Expected an error but received nil")
	}
}

func TestConfirm_WithUserLinkError(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo, _ := SharedSetup(t)

	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}
	expectedCharName := "Test Character Result"

	allianceId := int64(1)

	underlyingError := "I'm sorry, Dave. I'm afraid I can't do that."

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(
			&model.Character{
				CharacterId:   1,
				CharacterName: expectedCharName,
				CorporationId: 1,
				Corporation: model.Corporation{
					CorporationId:     1,
					CorporationName:   "Test Corporation Result",
					CorporationTicker: "TSTC",
					AllianceId:        &allianceId,
					Alliance: model.Alliance{
						AllianceId:     1,
						AllianceName:   "Test Alliance Result",
						AllianceTicker: "TSTA",
					},
				},
			},
		),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil),
		mockUserRepo.EXPECT().Save(
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(authConfirmRequest.AuthenticationCode,
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		).Return(errors.New(underlyingError)),
		mockAcceRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return([]string{"ROLE1", "ROLE2", "ROLE3"}, nil).Times(0),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	expectedError := "Error linking user: " + underlyingError

	if err != nil && err.Error() != expectedError {
		t.Errorf("Expected error message (%s) but received (%s)", expectedError, err)
	} else if err == nil {
		t.Error("Expected an error but received nil")
	}
}

func TestConfirm_WithRoleFindError(t *testing.T) {
	mockCtrl, _, mockUserRepo, mockCharRepo, _, _, mockAcceRepo, _ := SharedSetup(t)

	defer mockCtrl.Finish()

	var ctx context.Context
	var authConfirmResponse proto.AuthConfirmResponse
	authConfirmRequest := proto.AuthConfirmRequest{
		UserId:             "1234567890",
		AuthenticationCode: "123456789012",
	}

	authHandler := AuthHandler{}
	expectedCharName := "Test Character Result"

	allianceId := int64(1)

	underlyingError := "I'm sorry, Dave. I'm afraid I can't do that."

	gomock.InOrder(
		mockCharRepo.EXPECT().FindByAutenticationCode(authConfirmRequest.AuthenticationCode).Return(
			&model.Character{
				CharacterId:   1,
				CharacterName: expectedCharName,
				CorporationId: 1,
				Corporation: model.Corporation{
					CorporationId:     1,
					CorporationName:   "Test Corporation Result",
					CorporationTicker: "TSTC",
					AllianceId:        &allianceId,
					Alliance: model.Alliance{
						AllianceId:     1,
						AllianceName:   "Test Alliance Result",
						AllianceTicker: "TSTA",
					},
				},
			},
		),
		mockUserRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil),
		mockUserRepo.EXPECT().Save(
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		),
		mockUserRepo.EXPECT().LinkCharacterToUserByAuthCode(authConfirmRequest.AuthenticationCode,
			&model.User{
				ChatId: authConfirmRequest.UserId,
			},
		),
		mockAcceRepo.EXPECT().FindByChatId(authConfirmRequest.UserId).Return(nil, errors.New(underlyingError)),
	)

	err := authHandler.Confirm(ctx, &authConfirmRequest, &authConfirmResponse)

	expectedError := "Error finding roles: " + underlyingError

	if err != nil && err.Error() != expectedError {
		t.Errorf("Expected error message (%s) but received (%s)", expectedError, err)
	} else if err == nil {
		t.Error("Expected an error but received nil")
	}
}

func TestGetRoles(t *testing.T) {
	mockCtrl, _, mockUserRepo, _, _, _, mockAcceRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authHandler := AuthHandler{}

	underlyingError := "I'm sorry, Dave. I'm afraid I can't do that."

	gomock.InOrder(
		mockUserRepo.EXPECT().FindByChatId("1234567890").Return(&model.User{UserId: 1, ChatId: "1234567890"}),
		mockAcceRepo.EXPECT().FindByChatId("1234567890").Return(nil, errors.New(underlyingError)),
	)

	response := proto.AuthConfirmResponse{}

	err := authHandler.GetRoles(context.Background(), &proto.GetRolesRequest{UserId: "1234567890"}, &response)

	expectedError := "Error finding roles: " + underlyingError

	if err != nil && err.Error() != expectedError {
		t.Errorf("Expected error message (%s) but received (%s)", expectedError, err)
	} else if err == nil {
		t.Error("Expected an error but received nil")
	}
}

func TestGetRoles_WithRoleFindError(t *testing.T) {
	mockCtrl, _, mockUserRepo, _, _, _, mockAcceRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockUserRepo.EXPECT().FindByChatId("1234567890").Return(&model.User{UserId: 1, ChatId: "1234567890"}),
		mockAcceRepo.EXPECT().FindByChatId("1234567890").Return([]string{"ROLE1", "ROLE2"}, nil),
	)

	response := proto.AuthConfirmResponse{}

	err := authHandler.GetRoles(context.Background(), &proto.GetRolesRequest{UserId: "1234567890"}, &response)

	if err != nil {
		t.Fatal("Received an error when none were expected.")
	}

	if !response.Success {
		t.Fatal("Received false success when true was expected")
	}

	if len(response.CharacterName) != 0 {
		t.Fatalf("Received: (%s) for character name when empty was expected", response.CharacterName)
	}

	if len(response.Roles) != 2 {
		t.Fatalf("Expected 2 response but %d were received", len(response.Roles))
	}
}

func TestGetRolesNoUser(t *testing.T) {
	mockCtrl, _, mockUserRepo, _, _, _, mockAcceRepo, _ := SharedSetup(t)
	defer mockCtrl.Finish()

	authHandler := AuthHandler{}

	gomock.InOrder(
		mockUserRepo.EXPECT().FindByChatId("1234567890").Return(nil),
		mockAcceRepo.EXPECT().FindByChatId("1234567890").Return([]string{"ROLE1", "ROLE2"}, nil).Times(0),
	)

	response := proto.AuthConfirmResponse{}

	err := authHandler.GetRoles(context.Background(), &proto.GetRolesRequest{UserId: "1234567890"}, &response)

	if err != nil {
		t.Fatal("Received an error when none were expected.")
	}

	if response.Success {
		t.Fatal("Received true success when false was expected")
	}

	if len(response.Roles) > 0 {
		t.Fatal("Received some roles when we should have errored out")
	}
}

//TODO: Create test case for AuthWithUserAndChar?  Should overwrite or create new auth record?  DB Model supports multiple auths per character... by why?  Was a drunk?
