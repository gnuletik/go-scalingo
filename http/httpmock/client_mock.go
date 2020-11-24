// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo/http (interfaces: Client)

// Package httpmock is a generated GoMock package.
package httpmock

import (
	http0 "net/http"
	reflect "reflect"

	http "github.com/Scalingo/go-scalingo/v4/http"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BaseURL mocks base method.
func (m *MockClient) BaseURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURL indicates an expected call of BaseURL.
func (mr *MockClientMockRecorder) BaseURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURL", reflect.TypeOf((*MockClient)(nil).BaseURL))
}

// Do mocks base method.
func (m *MockClient) Do(arg0 *http.APIRequest) (*http0.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http0.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockClientMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClient)(nil).Do), arg0)
}

// DoRequest mocks base method.
func (m *MockClient) DoRequest(arg0 *http.APIRequest, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoRequest indicates an expected call of DoRequest.
func (mr *MockClientMockRecorder) DoRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoRequest", reflect.TypeOf((*MockClient)(nil).DoRequest), arg0, arg1)
}

// HTTPClient mocks base method.
func (m *MockClient) HTTPClient() *http0.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*http0.Client)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}

// IsAuthenticatedClient mocks base method.
func (m *MockClient) IsAuthenticatedClient() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthenticatedClient")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAuthenticatedClient indicates an expected call of IsAuthenticatedClient.
func (mr *MockClientMockRecorder) IsAuthenticatedClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthenticatedClient", reflect.TypeOf((*MockClient)(nil).IsAuthenticatedClient))
}

// ResourceAdd mocks base method.
func (m *MockClient) ResourceAdd(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceAdd", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceAdd indicates an expected call of ResourceAdd.
func (mr *MockClientMockRecorder) ResourceAdd(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceAdd", reflect.TypeOf((*MockClient)(nil).ResourceAdd), arg0, arg1, arg2)
}

// ResourceDelete mocks base method.
func (m *MockClient) ResourceDelete(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceDelete indicates an expected call of ResourceDelete.
func (mr *MockClientMockRecorder) ResourceDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceDelete", reflect.TypeOf((*MockClient)(nil).ResourceDelete), arg0, arg1)
}

// ResourceGet mocks base method.
func (m *MockClient) ResourceGet(arg0, arg1 string, arg2, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceGet indicates an expected call of ResourceGet.
func (mr *MockClientMockRecorder) ResourceGet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGet", reflect.TypeOf((*MockClient)(nil).ResourceGet), arg0, arg1, arg2, arg3)
}

// ResourceList mocks base method.
func (m *MockClient) ResourceList(arg0 string, arg1, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceList", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceList indicates an expected call of ResourceList.
func (mr *MockClientMockRecorder) ResourceList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceList", reflect.TypeOf((*MockClient)(nil).ResourceList), arg0, arg1, arg2)
}

// ResourceUpdate mocks base method.
func (m *MockClient) ResourceUpdate(arg0, arg1 string, arg2, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResourceUpdate indicates an expected call of ResourceUpdate.
func (mr *MockClientMockRecorder) ResourceUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceUpdate", reflect.TypeOf((*MockClient)(nil).ResourceUpdate), arg0, arg1, arg2, arg3)
}

// SubresourceAdd mocks base method.
func (m *MockClient) SubresourceAdd(arg0, arg1, arg2 string, arg3, arg4 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubresourceAdd", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubresourceAdd indicates an expected call of SubresourceAdd.
func (mr *MockClientMockRecorder) SubresourceAdd(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubresourceAdd", reflect.TypeOf((*MockClient)(nil).SubresourceAdd), arg0, arg1, arg2, arg3, arg4)
}

// SubresourceDelete mocks base method.
func (m *MockClient) SubresourceDelete(arg0, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubresourceDelete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubresourceDelete indicates an expected call of SubresourceDelete.
func (mr *MockClientMockRecorder) SubresourceDelete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubresourceDelete", reflect.TypeOf((*MockClient)(nil).SubresourceDelete), arg0, arg1, arg2, arg3)
}

// SubresourceGet mocks base method.
func (m *MockClient) SubresourceGet(arg0, arg1, arg2, arg3 string, arg4, arg5 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubresourceGet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubresourceGet indicates an expected call of SubresourceGet.
func (mr *MockClientMockRecorder) SubresourceGet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubresourceGet", reflect.TypeOf((*MockClient)(nil).SubresourceGet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// SubresourceList mocks base method.
func (m *MockClient) SubresourceList(arg0, arg1, arg2 string, arg3, arg4 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubresourceList", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubresourceList indicates an expected call of SubresourceList.
func (mr *MockClientMockRecorder) SubresourceList(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubresourceList", reflect.TypeOf((*MockClient)(nil).SubresourceList), arg0, arg1, arg2, arg3, arg4)
}

// SubresourceUpdate mocks base method.
func (m *MockClient) SubresourceUpdate(arg0, arg1, arg2, arg3 string, arg4, arg5 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubresourceUpdate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubresourceUpdate indicates an expected call of SubresourceUpdate.
func (mr *MockClientMockRecorder) SubresourceUpdate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubresourceUpdate", reflect.TypeOf((*MockClient)(nil).SubresourceUpdate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokenGenerator mocks base method.
func (m *MockClient) TokenGenerator() http.TokenGenerator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenGenerator")
	ret0, _ := ret[0].(http.TokenGenerator)
	return ret0
}

// TokenGenerator indicates an expected call of TokenGenerator.
func (mr *MockClientMockRecorder) TokenGenerator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenGenerator", reflect.TypeOf((*MockClient)(nil).TokenGenerator))
}
