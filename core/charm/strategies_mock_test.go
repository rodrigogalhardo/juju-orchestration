// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: StateCharm,State,StoreCharm,Store,JujuVersionValidator)

// Package charm is a generated GoMock package.
package charm

import (
	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v9"
	reflect "reflect"
)

// MockStateCharm is a mock of StateCharm interface
type MockStateCharm struct {
	ctrl     *gomock.Controller
	recorder *MockStateCharmMockRecorder
}

// MockStateCharmMockRecorder is the mock recorder for MockStateCharm
type MockStateCharmMockRecorder struct {
	mock *MockStateCharm
}

// NewMockStateCharm creates a new mock instance
func NewMockStateCharm(ctrl *gomock.Controller) *MockStateCharm {
	mock := &MockStateCharm{ctrl: ctrl}
	mock.recorder = &MockStateCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateCharm) EXPECT() *MockStateCharmMockRecorder {
	return m.recorder
}

// IsUploaded mocks base method
func (m *MockStateCharm) IsUploaded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUploaded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUploaded indicates an expected call of IsUploaded
func (mr *MockStateCharmMockRecorder) IsUploaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUploaded", reflect.TypeOf((*MockStateCharm)(nil).IsUploaded))
}

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// PrepareCharmUpload mocks base method
func (m *MockState) PrepareCharmUpload(arg0 *charm.URL) (StateCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(StateCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload
func (mr *MockStateMockRecorder) PrepareCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockState)(nil).PrepareCharmUpload), arg0)
}

// MockStoreCharm is a mock of StoreCharm interface
type MockStoreCharm struct {
	ctrl     *gomock.Controller
	recorder *MockStoreCharmMockRecorder
}

// MockStoreCharmMockRecorder is the mock recorder for MockStoreCharm
type MockStoreCharmMockRecorder struct {
	mock *MockStoreCharm
}

// NewMockStoreCharm creates a new mock instance
func NewMockStoreCharm(ctrl *gomock.Controller) *MockStoreCharm {
	mock := &MockStoreCharm{ctrl: ctrl}
	mock.recorder = &MockStoreCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreCharm) EXPECT() *MockStoreCharmMockRecorder {
	return m.recorder
}

// Actions mocks base method
func (m *MockStoreCharm) Actions() *charm.Actions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(*charm.Actions)
	return ret0
}

// Actions indicates an expected call of Actions
func (mr *MockStoreCharmMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockStoreCharm)(nil).Actions))
}

// Config mocks base method
func (m *MockStoreCharm) Config() *charm.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*charm.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockStoreCharmMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockStoreCharm)(nil).Config))
}

// LXDProfile mocks base method
func (m *MockStoreCharm) LXDProfile() *charm.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile
func (mr *MockStoreCharmMockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockStoreCharm)(nil).LXDProfile))
}

// Meta mocks base method
func (m *MockStoreCharm) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockStoreCharmMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockStoreCharm)(nil).Meta))
}

// Metrics mocks base method
func (m *MockStoreCharm) Metrics() *charm.Metrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(*charm.Metrics)
	return ret0
}

// Metrics indicates an expected call of Metrics
func (mr *MockStoreCharmMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockStoreCharm)(nil).Metrics))
}

// Revision mocks base method
func (m *MockStoreCharm) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision
func (mr *MockStoreCharmMockRecorder) Revision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockStoreCharm)(nil).Revision))
}

// Version mocks base method
func (m *MockStoreCharm) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockStoreCharmMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockStoreCharm)(nil).Version))
}

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockStore) Download(arg0 *charm.URL, arg1 string, arg2 Origin) (StoreCharm, ChecksumCheckFn, Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1, arg2)
	ret0, _ := ret[0].(StoreCharm)
	ret1, _ := ret[1].(ChecksumCheckFn)
	ret2, _ := ret[2].(Origin)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Download indicates an expected call of Download
func (mr *MockStoreMockRecorder) Download(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockStore)(nil).Download), arg0, arg1, arg2)
}

// DownloadOrigin mocks base method
func (m *MockStore) DownloadOrigin(arg0 *charm.URL, arg1 Origin) (Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadOrigin", arg0, arg1)
	ret0, _ := ret[0].(Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadOrigin indicates an expected call of DownloadOrigin
func (mr *MockStoreMockRecorder) DownloadOrigin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadOrigin", reflect.TypeOf((*MockStore)(nil).DownloadOrigin), arg0, arg1)
}

// Validate mocks base method
func (m *MockStore) Validate(arg0 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockStoreMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockStore)(nil).Validate), arg0)
}

// MockJujuVersionValidator is a mock of JujuVersionValidator interface
type MockJujuVersionValidator struct {
	ctrl     *gomock.Controller
	recorder *MockJujuVersionValidatorMockRecorder
}

// MockJujuVersionValidatorMockRecorder is the mock recorder for MockJujuVersionValidator
type MockJujuVersionValidatorMockRecorder struct {
	mock *MockJujuVersionValidator
}

// NewMockJujuVersionValidator creates a new mock instance
func NewMockJujuVersionValidator(ctrl *gomock.Controller) *MockJujuVersionValidator {
	mock := &MockJujuVersionValidator{ctrl: ctrl}
	mock.recorder = &MockJujuVersionValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJujuVersionValidator) EXPECT() *MockJujuVersionValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockJujuVersionValidator) Validate(arg0 *charm.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockJujuVersionValidatorMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockJujuVersionValidator)(nil).Validate), arg0)
}
