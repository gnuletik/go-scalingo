// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/v4 (interfaces: BackupsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	context "context"
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockBackupsService is a mock of BackupsService interface.
type MockBackupsService struct {
	ctrl     *gomock.Controller
	recorder *MockBackupsServiceMockRecorder
}

// MockBackupsServiceMockRecorder is the mock recorder for MockBackupsService.
type MockBackupsServiceMockRecorder struct {
	mock *MockBackupsService
}

// NewMockBackupsService creates a new mock instance.
func NewMockBackupsService(ctrl *gomock.Controller) *MockBackupsService {
	mock := &MockBackupsService{ctrl: ctrl}
	mock.recorder = &MockBackupsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupsService) EXPECT() *MockBackupsServiceMockRecorder {
	return m.recorder
}

// BackupCreate mocks base method.
func (m *MockBackupsService) BackupCreate(arg0 context.Context, arg1, arg2 string) (*scalingo.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].(*scalingo.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupCreate indicates an expected call of BackupCreate.
func (mr *MockBackupsServiceMockRecorder) BackupCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupCreate", reflect.TypeOf((*MockBackupsService)(nil).BackupCreate), arg0, arg1, arg2)
}

// BackupDownloadURL mocks base method.
func (m *MockBackupsService) BackupDownloadURL(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupDownloadURL", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupDownloadURL indicates an expected call of BackupDownloadURL.
func (mr *MockBackupsServiceMockRecorder) BackupDownloadURL(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupDownloadURL", reflect.TypeOf((*MockBackupsService)(nil).BackupDownloadURL), arg0, arg1, arg2, arg3)
}

// BackupList mocks base method.
func (m *MockBackupsService) BackupList(arg0 context.Context, arg1, arg2 string) ([]scalingo.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]scalingo.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupList indicates an expected call of BackupList.
func (mr *MockBackupsServiceMockRecorder) BackupList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupList", reflect.TypeOf((*MockBackupsService)(nil).BackupList), arg0, arg1, arg2)
}

// BackupShow mocks base method.
func (m *MockBackupsService) BackupShow(arg0 context.Context, arg1, arg2, arg3 string) (*scalingo.Backup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackupShow", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*scalingo.Backup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackupShow indicates an expected call of BackupShow.
func (mr *MockBackupsServiceMockRecorder) BackupShow(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackupShow", reflect.TypeOf((*MockBackupsService)(nil).BackupShow), arg0, arg1, arg2, arg3)
}
