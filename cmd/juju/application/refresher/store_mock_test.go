// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/store (interfaces: CharmAdder)

// Package refresher is a generated GoMock package.
package refresher

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v11"
	charm0 "github.com/juju/juju/api/common/charm"
)

// MockCharmAdder is a mock of CharmAdder interface.
type MockCharmAdder struct {
	ctrl     *gomock.Controller
	recorder *MockCharmAdderMockRecorder
}

// MockCharmAdderMockRecorder is the mock recorder for MockCharmAdder.
type MockCharmAdderMockRecorder struct {
	mock *MockCharmAdder
}

// NewMockCharmAdder creates a new mock instance.
func NewMockCharmAdder(ctrl *gomock.Controller) *MockCharmAdder {
	mock := &MockCharmAdder{ctrl: ctrl}
	mock.recorder = &MockCharmAdderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmAdder) EXPECT() *MockCharmAdderMockRecorder {
	return m.recorder
}

// AddCharm mocks base method.
func (m *MockCharmAdder) AddCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharm indicates an expected call of AddCharm.
func (mr *MockCharmAdderMockRecorder) AddCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddCharm), arg0, arg1, arg2)
}

// AddLocalCharm mocks base method.
func (m *MockCharmAdder) AddLocalCharm(arg0 *charm.URL, arg1 charm.Charm, arg2 bool) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalCharm indicates an expected call of AddLocalCharm.
func (mr *MockCharmAdderMockRecorder) AddLocalCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalCharm", reflect.TypeOf((*MockCharmAdder)(nil).AddLocalCharm), arg0, arg1, arg2)
}

// CheckCharmPlacement mocks base method.
func (m *MockCharmAdder) CheckCharmPlacement(arg0 string, arg1 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCharmPlacement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCharmPlacement indicates an expected call of CheckCharmPlacement.
func (mr *MockCharmAdderMockRecorder) CheckCharmPlacement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCharmPlacement", reflect.TypeOf((*MockCharmAdder)(nil).CheckCharmPlacement), arg0, arg1)
}
