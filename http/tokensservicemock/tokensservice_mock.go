// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/http (interfaces: TokensService)

// Package tokensservicemock is a generated GoMock package.
package tokensservicemock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTokensService is a mock of TokensService interface
type MockTokensService struct {
	ctrl     *gomock.Controller
	recorder *MockTokensServiceMockRecorder
}

// MockTokensServiceMockRecorder is the mock recorder for MockTokensService
type MockTokensServiceMockRecorder struct {
	mock *MockTokensService
}

// NewMockTokensService creates a new mock instance
func NewMockTokensService(ctrl *gomock.Controller) *MockTokensService {
	mock := &MockTokensService{ctrl: ctrl}
	mock.recorder = &MockTokensServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTokensService) EXPECT() *MockTokensServiceMockRecorder {
	return m.recorder
}

// TokenExchange mocks base method
func (m *MockTokensService) TokenExchange(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenExchange", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenExchange indicates an expected call of TokenExchange
func (mr *MockTokensServiceMockRecorder) TokenExchange(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenExchange", reflect.TypeOf((*MockTokensService)(nil).TokenExchange), arg0)
}