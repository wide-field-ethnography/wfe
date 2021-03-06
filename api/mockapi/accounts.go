// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/citwild/wfe/api (interfaces: AccountsServer)

package mockapi

import (
	api "github.com/citwild/wfe/api"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of AccountsServer interface
type MockAccountsServer struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountsServerRecorder
}

// Recorder for MockAccountsServer (not exported)
type _MockAccountsServerRecorder struct {
	mock *MockAccountsServer
}

func NewMockAccountsServer(ctrl *gomock.Controller) *MockAccountsServer {
	mock := &MockAccountsServer{ctrl: ctrl}
	mock.recorder = &_MockAccountsServerRecorder{mock}
	return mock
}

func (_m *MockAccountsServer) EXPECT() *_MockAccountsServerRecorder {
	return _m.recorder
}

func (_m *MockAccountsServer) Create(_param0 context.Context, _param1 *api.NewAccount) (*api.CreatedAccount, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0, _param1)
	ret0, _ := ret[0].(*api.CreatedAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccountsServerRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1)
}
