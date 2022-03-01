// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/v4 (interfaces: OperationsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockOperationsService is a mock of OperationsService interface
type MockOperationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsServiceMockRecorder
}

// MockOperationsServiceMockRecorder is the mock recorder for MockOperationsService
type MockOperationsServiceMockRecorder struct {
	mock *MockOperationsService
}

// NewMockOperationsService creates a new mock instance
func NewMockOperationsService(ctrl *gomock.Controller) *MockOperationsService {
	mock := &MockOperationsService{ctrl: ctrl}
	mock.recorder = &MockOperationsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationsService) EXPECT() *MockOperationsServiceMockRecorder {
	return m.recorder
}

// OperationsShow mocks base method
func (m *MockOperationsService) OperationsShow(arg0, arg1 string) (*scalingo.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationsShow", arg0, arg1)
	ret0, _ := ret[0].(*scalingo.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperationsShow indicates an expected call of OperationsShow
func (mr *MockOperationsServiceMockRecorder) OperationsShow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationsShow", reflect.TypeOf((*MockOperationsService)(nil).OperationsShow), arg0, arg1)
}