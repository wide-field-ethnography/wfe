// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../users.go

package mockstore

import (
	gomock "github.com/golang/mock/gomock"
	wfe "github.com/citwild/wfe/api/wfe"
	context "golang.org/x/net/context"
)

// Mock of Accounts interface
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountsRecorder
}

// Recorder for MockAccounts (not exported)
type _MockAccountsRecorder struct {
	mock *MockAccounts
}

func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &_MockAccountsRecorder{mock}
	return mock
}

func (_m *MockAccounts) EXPECT() *_MockAccountsRecorder {
	return _m.recorder
}

func (_m *MockAccounts) Create(ctx context.Context, newUser *wfe.User, email *wfe.EmailAddr) (*wfe.User, error) {
	ret := _m.ctrl.Call(_m, "Create", ctx, newUser, email)
	ret0, _ := ret[0].(*wfe.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAccountsRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0, arg1, arg2)
}
