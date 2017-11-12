// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../repository/accesses.go

package mocks

import (
	model "github.com/abaeve/auth-srv/model"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AccessesRepository interface
type MockAccessesRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAccessesRepositoryRecorder
}

// Recorder for MockAccessesRepository (not exported)
type _MockAccessesRepositoryRecorder struct {
	mock *MockAccessesRepository
}

func NewMockAccessesRepository(ctrl *gomock.Controller) *MockAccessesRepository {
	mock := &MockAccessesRepository{ctrl: ctrl}
	mock.recorder = &_MockAccessesRepositoryRecorder{mock}
	return mock
}

func (_m *MockAccessesRepository) EXPECT() *_MockAccessesRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAccessesRepository) SaveAllianceAndCorpRole(allianceId int64, corporationId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceAndCorpRole", allianceId, corporationId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceAndCorpRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceAndCorpRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveAllianceRole(allianceId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceRole", allianceId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCorporationRole(corporationId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationRole", corporationId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCharacterRole(characterId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCharacterRole", characterId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCharacterRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCharacterRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveAllianceCharacterLeadershipRole(allianceId int64, characterId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceCharacterLeadershipRole", allianceId, characterId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveCorporationCharacterLeadershipRole(corporationId int64, characterId int64, role *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationCharacterLeadershipRole", corporationId, characterId, role)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteAllianceAndCorpRole(allianceId int64, corporationId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceAndCorpRole", allianceId, corporationId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceAndCorpRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceAndCorpRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteAllianceRole(allianceId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceRole", allianceId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceRole", arg0, arg1)
}

func (_m *MockAccessesRepository) DeleteCorporationRole(corporationId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCorporationRole", corporationId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCorporationRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCorporationRole", arg0, arg1)
}

func (_m *MockAccessesRepository) DeleteCharacterRole(characterId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCharacterRole", characterId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCharacterRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCharacterRole", arg0, arg1)
}

func (_m *MockAccessesRepository) DeleteAllianceCharacterLeadershipRole(allianceId int64, characterId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceCharacterLeadershipRole", allianceId, characterId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteCorporationCharacterLeadershipRole(corporationId int64, characterId int64, role *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCorporationCharacterLeadershipRole", corporationId, characterId, role)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCorporationCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCorporationCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) FindByChatId(chatId string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "FindByChatId", chatId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) FindByChatId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByChatId", arg0)
}