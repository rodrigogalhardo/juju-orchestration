// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/externalcontrollerupdater (interfaces: EcService)

// Package externalcontrollerupdater_test is a generated GoMock package.
package externalcontrollerupdater_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	crossmodel "github.com/juju/juju/core/crossmodel"
)

// MockEcService is a mock of EcService interface.
type MockEcService struct {
	ctrl     *gomock.Controller
	recorder *MockEcServiceMockRecorder
}

// MockEcServiceMockRecorder is the mock recorder for MockEcService.
type MockEcServiceMockRecorder struct {
	mock *MockEcService
}

// NewMockEcService creates a new mock instance.
func NewMockEcService(ctrl *gomock.Controller) *MockEcService {
	mock := &MockEcService{ctrl: ctrl}
	mock.recorder = &MockEcServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcService) EXPECT() *MockEcServiceMockRecorder {
	return m.recorder
}

// UpdateExternalController mocks base method.
func (m *MockEcService) UpdateExternalController(arg0 context.Context, arg1 crossmodel.ControllerInfo, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateExternalController", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExternalController indicates an expected call of UpdateExternalController.
func (mr *MockEcServiceMockRecorder) UpdateExternalController(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExternalController", reflect.TypeOf((*MockEcService)(nil).UpdateExternalController), varargs...)
}
