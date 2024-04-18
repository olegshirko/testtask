// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"
	entity "testTask/internal/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockAsset is a mock of Asset interface.
type MockAsset struct {
	ctrl     *gomock.Controller
	recorder *MockAssetMockRecorder
}

// MockAssetMockRecorder is the mock recorder for MockAsset.
type MockAssetMockRecorder struct {
	mock *MockAsset
}

// NewMockAsset creates a new mock instance.
func NewMockAsset(ctrl *gomock.Controller) *MockAsset {
	mock := &MockAsset{ctrl: ctrl}
	mock.recorder = &MockAssetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsset) EXPECT() *MockAssetMockRecorder {
	return m.recorder
}

// DropAsset mocks base method.
func (m *MockAsset) DropAsset(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropAsset", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DropAsset indicates an expected call of DropAsset.
func (mr *MockAssetMockRecorder) DropAsset(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropAsset", reflect.TypeOf((*MockAsset)(nil).DropAsset), arg0, arg1, arg2)
}

// GetUserIdByToken mocks base method.
func (m *MockAsset) GetUserIdByToken(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByToken", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByToken indicates an expected call of GetUserIdByToken.
func (mr *MockAssetMockRecorder) GetUserIdByToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByToken", reflect.TypeOf((*MockAsset)(nil).GetUserIdByToken), arg0, arg1)
}

// History mocks base method.
func (m *MockAsset) History(arg0 context.Context, arg1 int64) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "History", arg0, arg1)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// History indicates an expected call of History.
func (mr *MockAssetMockRecorder) History(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "History", reflect.TypeOf((*MockAsset)(nil).History), arg0, arg1)
}

// Session mocks base method.
func (m *MockAsset) Session(arg0 context.Context, arg1 string, arg2 entity.AuthData) (entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Session indicates an expected call of Session.
func (mr *MockAssetMockRecorder) Session(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockAsset)(nil).Session), arg0, arg1, arg2)
}

// UploadAsset mocks base method.
func (m *MockAsset) UploadAsset(arg0 context.Context, arg1 []byte, arg2 string, arg3 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAsset", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAsset indicates an expected call of UploadAsset.
func (mr *MockAssetMockRecorder) UploadAsset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAsset", reflect.TypeOf((*MockAsset)(nil).UploadAsset), arg0, arg1, arg2, arg3)
}

// MockAssetRepo is a mock of AssetRepo interface.
type MockAssetRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAssetRepoMockRecorder
}

// MockAssetRepoMockRecorder is the mock recorder for MockAssetRepo.
type MockAssetRepoMockRecorder struct {
	mock *MockAssetRepo
}

// NewMockAssetRepo creates a new mock instance.
func NewMockAssetRepo(ctrl *gomock.Controller) *MockAssetRepo {
	mock := &MockAssetRepo{ctrl: ctrl}
	mock.recorder = &MockAssetRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssetRepo) EXPECT() *MockAssetRepoMockRecorder {
	return m.recorder
}

// DeleteAsset mocks base method.
func (m *MockAssetRepo) DeleteAsset(arg0 context.Context, arg1 string, arg2 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAsset", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAsset indicates an expected call of DeleteAsset.
func (mr *MockAssetRepoMockRecorder) DeleteAsset(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAsset", reflect.TypeOf((*MockAssetRepo)(nil).DeleteAsset), arg0, arg1, arg2)
}

// GetHistory mocks base method.
func (m *MockAssetRepo) GetHistory(arg0 context.Context, arg1 int64) ([]entity.Asset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", arg0, arg1)
	ret0, _ := ret[0].([]entity.Asset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockAssetRepoMockRecorder) GetHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockAssetRepo)(nil).GetHistory), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockAssetRepo) GetSession(arg0 context.Context, arg1 entity.User) (entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockAssetRepoMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockAssetRepo)(nil).GetSession), arg0, arg1)
}

// GetUserByAuthData mocks base method.
func (m *MockAssetRepo) GetUserByAuthData(arg0 context.Context, arg1 entity.AuthData) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuthData", arg0, arg1)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuthData indicates an expected call of GetUserByAuthData.
func (mr *MockAssetRepoMockRecorder) GetUserByAuthData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuthData", reflect.TypeOf((*MockAssetRepo)(nil).GetUserByAuthData), arg0, arg1)
}

// GetUserIdByToken mocks base method.
func (m *MockAssetRepo) GetUserIdByToken(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdByToken", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdByToken indicates an expected call of GetUserIdByToken.
func (mr *MockAssetRepoMockRecorder) GetUserIdByToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdByToken", reflect.TypeOf((*MockAssetRepo)(nil).GetUserIdByToken), arg0, arg1)
}

// StoreAsset mocks base method.
func (m *MockAssetRepo) StoreAsset(arg0 context.Context, arg1 []byte, arg2 string, arg3 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAsset", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreAsset indicates an expected call of StoreAsset.
func (mr *MockAssetRepoMockRecorder) StoreAsset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAsset", reflect.TypeOf((*MockAssetRepo)(nil).StoreAsset), arg0, arg1, arg2, arg3)
}

// StoreSession mocks base method.
func (m *MockAssetRepo) StoreSession(arg0 context.Context, arg1 entity.User, arg2 string) (entity.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(entity.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreSession indicates an expected call of StoreSession.
func (mr *MockAssetRepoMockRecorder) StoreSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSession", reflect.TypeOf((*MockAssetRepo)(nil).StoreSession), arg0, arg1, arg2)
}
