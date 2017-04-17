// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/abaeve/auth-srv/repository (interfaces: AccessesRepository,AllianceRepository,CorporationRepository,CharacterRepository,UserRepository,RoleRepository,AuthenticationCodeRepository)

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

func (_m *MockAccessesRepository) DeleteAllianceAndCorpRole(_param0 int64, _param1 int64, _param2 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceAndCorpRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceAndCorpRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceAndCorpRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteAllianceCharacterLeadershipRole(_param0 int64, _param1 int64, _param2 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceCharacterLeadershipRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteAllianceRole(_param0 int64, _param1 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteAllianceRole", _param0, _param1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteAllianceRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAllianceRole", arg0, arg1)
}

func (_m *MockAccessesRepository) DeleteCharacterRole(_param0 int64, _param1 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCharacterRole", _param0, _param1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCharacterRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCharacterRole", arg0, arg1)
}

func (_m *MockAccessesRepository) DeleteCorporationCharacterLeadershipRole(_param0 int64, _param1 int64, _param2 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCorporationCharacterLeadershipRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCorporationCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCorporationCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) DeleteCorporationRole(_param0 int64, _param1 *model.Role) (int64, error) {
	ret := _m.ctrl.Call(_m, "DeleteCorporationRole", _param0, _param1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) DeleteCorporationRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCorporationRole", arg0, arg1)
}

func (_m *MockAccessesRepository) FindByChatId(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "FindByChatId", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccessesRepositoryRecorder) FindByChatId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByChatId", arg0)
}

func (_m *MockAccessesRepository) SaveAllianceAndCorpRole(_param0 int64, _param1 int64, _param2 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceAndCorpRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceAndCorpRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceAndCorpRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveAllianceCharacterLeadershipRole(_param0 int64, _param1 int64, _param2 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceCharacterLeadershipRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveAllianceRole(_param0 int64, _param1 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveAllianceRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveAllianceRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveAllianceRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCharacterRole(_param0 int64, _param1 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCharacterRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCharacterRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCharacterRole", arg0, arg1)
}

func (_m *MockAccessesRepository) SaveCorporationCharacterLeadershipRole(_param0 int64, _param1 int64, _param2 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationCharacterLeadershipRole", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationCharacterLeadershipRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationCharacterLeadershipRole", arg0, arg1, arg2)
}

func (_m *MockAccessesRepository) SaveCorporationRole(_param0 int64, _param1 *model.Role) error {
	ret := _m.ctrl.Call(_m, "SaveCorporationRole", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAccessesRepositoryRecorder) SaveCorporationRole(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveCorporationRole", arg0, arg1)
}

// Mock of AllianceRepository interface
type MockAllianceRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAllianceRepositoryRecorder
}

// Recorder for MockAllianceRepository (not exported)
type _MockAllianceRepositoryRecorder struct {
	mock *MockAllianceRepository
}

func NewMockAllianceRepository(ctrl *gomock.Controller) *MockAllianceRepository {
	mock := &MockAllianceRepository{ctrl: ctrl}
	mock.recorder = &_MockAllianceRepositoryRecorder{mock}
	return mock
}

func (_m *MockAllianceRepository) EXPECT() *_MockAllianceRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAllianceRepository) FindByAllianceId(_param0 int64) *model.Alliance {
	ret := _m.ctrl.Call(_m, "FindByAllianceId", _param0)
	ret0, _ := ret[0].(*model.Alliance)
	return ret0
}

func (_mr *_MockAllianceRepositoryRecorder) FindByAllianceId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByAllianceId", arg0)
}

func (_m *MockAllianceRepository) Save(_param0 *model.Alliance) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAllianceRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// Mock of CorporationRepository interface
type MockCorporationRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockCorporationRepositoryRecorder
}

// Recorder for MockCorporationRepository (not exported)
type _MockCorporationRepositoryRecorder struct {
	mock *MockCorporationRepository
}

func NewMockCorporationRepository(ctrl *gomock.Controller) *MockCorporationRepository {
	mock := &MockCorporationRepository{ctrl: ctrl}
	mock.recorder = &_MockCorporationRepositoryRecorder{mock}
	return mock
}

func (_m *MockCorporationRepository) EXPECT() *_MockCorporationRepositoryRecorder {
	return _m.recorder
}

func (_m *MockCorporationRepository) FindByCorporationId(_param0 int64) *model.Corporation {
	ret := _m.ctrl.Call(_m, "FindByCorporationId", _param0)
	ret0, _ := ret[0].(*model.Corporation)
	return ret0
}

func (_mr *_MockCorporationRepositoryRecorder) FindByCorporationId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByCorporationId", arg0)
}

func (_m *MockCorporationRepository) Save(_param0 *model.Corporation) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCorporationRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// Mock of CharacterRepository interface
type MockCharacterRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockCharacterRepositoryRecorder
}

// Recorder for MockCharacterRepository (not exported)
type _MockCharacterRepositoryRecorder struct {
	mock *MockCharacterRepository
}

func NewMockCharacterRepository(ctrl *gomock.Controller) *MockCharacterRepository {
	mock := &MockCharacterRepository{ctrl: ctrl}
	mock.recorder = &_MockCharacterRepositoryRecorder{mock}
	return mock
}

func (_m *MockCharacterRepository) EXPECT() *_MockCharacterRepositoryRecorder {
	return _m.recorder
}

func (_m *MockCharacterRepository) FindByAutenticationCode(_param0 string) *model.Character {
	ret := _m.ctrl.Call(_m, "FindByAutenticationCode", _param0)
	ret0, _ := ret[0].(*model.Character)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) FindByAutenticationCode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByAutenticationCode", arg0)
}

func (_m *MockCharacterRepository) FindByCharacterId(_param0 int64) *model.Character {
	ret := _m.ctrl.Call(_m, "FindByCharacterId", _param0)
	ret0, _ := ret[0].(*model.Character)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) FindByCharacterId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByCharacterId", arg0)
}

func (_m *MockCharacterRepository) Save(_param0 *model.Character) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCharacterRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// Mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockUserRepositoryRecorder
}

// Recorder for MockUserRepository (not exported)
type _MockUserRepositoryRecorder struct {
	mock *MockUserRepository
}

func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &_MockUserRepositoryRecorder{mock}
	return mock
}

func (_m *MockUserRepository) EXPECT() *_MockUserRepositoryRecorder {
	return _m.recorder
}

func (_m *MockUserRepository) FindByChatId(_param0 string) *model.User {
	ret := _m.ctrl.Call(_m, "FindByChatId", _param0)
	ret0, _ := ret[0].(*model.User)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) FindByChatId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByChatId", arg0)
}

func (_m *MockUserRepository) LinkCharacterToUserByAuthCode(_param0 string, _param1 *model.User) error {
	ret := _m.ctrl.Call(_m, "LinkCharacterToUserByAuthCode", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) LinkCharacterToUserByAuthCode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkCharacterToUserByAuthCode", arg0, arg1)
}

func (_m *MockUserRepository) Save(_param0 *model.User) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// Mock of RoleRepository interface
type MockRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockRoleRepositoryRecorder
}

// Recorder for MockRoleRepository (not exported)
type _MockRoleRepositoryRecorder struct {
	mock *MockRoleRepository
}

func NewMockRoleRepository(ctrl *gomock.Controller) *MockRoleRepository {
	mock := &MockRoleRepository{ctrl: ctrl}
	mock.recorder = &_MockRoleRepositoryRecorder{mock}
	return mock
}

func (_m *MockRoleRepository) EXPECT() *_MockRoleRepositoryRecorder {
	return _m.recorder
}

func (_m *MockRoleRepository) FindByRoleName(_param0 string) *model.Role {
	ret := _m.ctrl.Call(_m, "FindByRoleName", _param0)
	ret0, _ := ret[0].(*model.Role)
	return ret0
}

func (_mr *_MockRoleRepositoryRecorder) FindByRoleName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindByRoleName", arg0)
}

func (_m *MockRoleRepository) Save(_param0 *model.Role) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRoleRepositoryRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// Mock of AuthenticationCodeRepository interface
type MockAuthenticationCodeRepository struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthenticationCodeRepositoryRecorder
}

// Recorder for MockAuthenticationCodeRepository (not exported)
type _MockAuthenticationCodeRepositoryRecorder struct {
	mock *MockAuthenticationCodeRepository
}

func NewMockAuthenticationCodeRepository(ctrl *gomock.Controller) *MockAuthenticationCodeRepository {
	mock := &MockAuthenticationCodeRepository{ctrl: ctrl}
	mock.recorder = &_MockAuthenticationCodeRepositoryRecorder{mock}
	return mock
}

func (_m *MockAuthenticationCodeRepository) EXPECT() *_MockAuthenticationCodeRepositoryRecorder {
	return _m.recorder
}

func (_m *MockAuthenticationCodeRepository) Save(_param0 *model.Character, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Save", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthenticationCodeRepositoryRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0, arg1)
}
