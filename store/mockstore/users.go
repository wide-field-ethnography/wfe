// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/citwild/wfe/store (interfaces: UsersStore)

package mockstore

import (
	"github.com/citwild/wfe/api"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
)

// Mock of UsersStore interface
type MockUsersStore struct {
	ctrl     *gomock.Controller
	recorder *_MockUsersStoreRecorder
}

// Recorder for MockUsersStore (not exported)
type _MockUsersStoreRecorder struct {
	mock *MockUsersStore
}

func NewMockUsersStore(ctrl *gomock.Controller) *MockUsersStore {
	mock := &MockUsersStore{ctrl: ctrl}
	mock.recorder = &_MockUsersStoreRecorder{mock}
	return mock
}

func (_m *MockUsersStore) EXPECT() *_MockUsersStoreRecorder {
	return _m.recorder
}

func (_m *MockUsersStore) Get(_param0 context.Context, _param1 api.UserSpec) (*api.User, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(*api.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUsersStoreRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}
