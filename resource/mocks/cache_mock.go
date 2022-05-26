// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/resource (interfaces: Resources)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resource "github.com/juju/charm/v8/resource"
	resource0 "github.com/juju/juju/core/resources"
	state "github.com/juju/juju/state"
)

// MockResources is a mock of Resources interface.
type MockResources struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesMockRecorder
}

// MockResourcesMockRecorder is the mock recorder for MockResources.
type MockResourcesMockRecorder struct {
	mock *MockResources
}

// NewMockResources creates a new mock instance.
func NewMockResources(ctrl *gomock.Controller) *MockResources {
	mock := &MockResources{ctrl: ctrl}
	mock.recorder = &MockResourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResources) EXPECT() *MockResourcesMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockResources) GetResource(arg0, arg1 string) (resource0.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0, arg1)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockResourcesMockRecorder) GetResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockResources)(nil).GetResource), arg0, arg1)
}

// OpenResource mocks base method.
func (m *MockResources) OpenResource(arg0, arg1 string) (resource0.Resource, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenResource", arg0, arg1)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OpenResource indicates an expected call of OpenResource.
func (mr *MockResourcesMockRecorder) OpenResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenResource", reflect.TypeOf((*MockResources)(nil).OpenResource), arg0, arg1)
}

// OpenResourceForUniter mocks base method.
func (m *MockResources) OpenResourceForUniter(arg0, arg1 string) (resource0.Resource, io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenResourceForUniter", arg0, arg1)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(io.ReadCloser)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OpenResourceForUniter indicates an expected call of OpenResourceForUniter.
func (mr *MockResourcesMockRecorder) OpenResourceForUniter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenResourceForUniter", reflect.TypeOf((*MockResources)(nil).OpenResourceForUniter), arg0, arg1)
}

// SetResource mocks base method.
func (m *MockResources) SetResource(arg0, arg1 string, arg2 resource.Resource, arg3 io.Reader, arg4 state.IncrementCharmModifiedVersionType) (resource0.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetResource", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(resource0.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetResource indicates an expected call of SetResource.
func (mr *MockResourcesMockRecorder) SetResource(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResource", reflect.TypeOf((*MockResources)(nil).SetResource), arg0, arg1, arg2, arg3, arg4)
}
