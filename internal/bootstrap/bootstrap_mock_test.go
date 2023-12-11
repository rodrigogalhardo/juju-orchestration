// Code generated by MockGen. DO NOT EDIT.
// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	context "context"
	io "io"
	http "net/http"
	reflect "reflect"

	charm "github.com/juju/charm/v12"
	services "github.com/juju/juju/apiserver/facades/client/charms/services"
	base "github.com/juju/juju/core/base"
	charm0 "github.com/juju/juju/core/charm"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	objectstore "github.com/juju/juju/core/objectstore"
	config "github.com/juju/juju/environs/config"
	charmhub "github.com/juju/juju/internal/charmhub"
	state "github.com/juju/juju/state"
	binarystorage "github.com/juju/juju/state/binarystorage"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockAgentBinaryStorage is a mock of AgentBinaryStorage interface.
type MockAgentBinaryStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAgentBinaryStorageMockRecorder
}

// MockAgentBinaryStorageMockRecorder is the mock recorder for MockAgentBinaryStorage.
type MockAgentBinaryStorageMockRecorder struct {
	mock *MockAgentBinaryStorage
}

// NewMockAgentBinaryStorage creates a new mock instance.
func NewMockAgentBinaryStorage(ctrl *gomock.Controller) *MockAgentBinaryStorage {
	mock := &MockAgentBinaryStorage{ctrl: ctrl}
	mock.recorder = &MockAgentBinaryStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentBinaryStorage) EXPECT() *MockAgentBinaryStorageMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockAgentBinaryStorage) Add(arg0 context.Context, arg1 io.Reader, arg2 binarystorage.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockAgentBinaryStorageMockRecorder) Add(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockAgentBinaryStorage)(nil).Add), arg0, arg1, arg2)
}

// MockControllerCharmDeployer is a mock of ControllerCharmDeployer interface.
type MockControllerCharmDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockControllerCharmDeployerMockRecorder
}

// MockControllerCharmDeployerMockRecorder is the mock recorder for MockControllerCharmDeployer.
type MockControllerCharmDeployerMockRecorder struct {
	mock *MockControllerCharmDeployer
}

// NewMockControllerCharmDeployer creates a new mock instance.
func NewMockControllerCharmDeployer(ctrl *gomock.Controller) *MockControllerCharmDeployer {
	mock := &MockControllerCharmDeployer{ctrl: ctrl}
	mock.recorder = &MockControllerCharmDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerCharmDeployer) EXPECT() *MockControllerCharmDeployerMockRecorder {
	return m.recorder
}

// AddControllerApplication mocks base method.
func (m *MockControllerCharmDeployer) AddControllerApplication(arg0 context.Context, arg1 string, arg2 *charm0.Origin, arg3 string) (Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddControllerApplication", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddControllerApplication indicates an expected call of AddControllerApplication.
func (mr *MockControllerCharmDeployerMockRecorder) AddControllerApplication(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddControllerApplication", reflect.TypeOf((*MockControllerCharmDeployer)(nil).AddControllerApplication), arg0, arg1, arg2, arg3)
}

// CompleteProcess mocks base method.
func (m *MockControllerCharmDeployer) CompleteProcess(arg0 context.Context, arg1 Unit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteProcess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteProcess indicates an expected call of CompleteProcess.
func (mr *MockControllerCharmDeployerMockRecorder) CompleteProcess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteProcess", reflect.TypeOf((*MockControllerCharmDeployer)(nil).CompleteProcess), arg0, arg1)
}

// ControllerAddress mocks base method.
func (m *MockControllerCharmDeployer) ControllerAddress(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerAddress", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerAddress indicates an expected call of ControllerAddress.
func (mr *MockControllerCharmDeployerMockRecorder) ControllerAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerAddress", reflect.TypeOf((*MockControllerCharmDeployer)(nil).ControllerAddress), arg0)
}

// ControllerCharmArch mocks base method.
func (m *MockControllerCharmDeployer) ControllerCharmArch() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerCharmArch")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerCharmArch indicates an expected call of ControllerCharmArch.
func (mr *MockControllerCharmDeployerMockRecorder) ControllerCharmArch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerCharmArch", reflect.TypeOf((*MockControllerCharmDeployer)(nil).ControllerCharmArch))
}

// ControllerCharmBase mocks base method.
func (m *MockControllerCharmDeployer) ControllerCharmBase() (base.Base, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerCharmBase")
	ret0, _ := ret[0].(base.Base)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerCharmBase indicates an expected call of ControllerCharmBase.
func (mr *MockControllerCharmDeployerMockRecorder) ControllerCharmBase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerCharmBase", reflect.TypeOf((*MockControllerCharmDeployer)(nil).ControllerCharmBase))
}

// DeployCharmhubCharm mocks base method.
func (m *MockControllerCharmDeployer) DeployCharmhubCharm(arg0 context.Context, arg1 string, arg2 base.Base) (string, *charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployCharmhubCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeployCharmhubCharm indicates an expected call of DeployCharmhubCharm.
func (mr *MockControllerCharmDeployerMockRecorder) DeployCharmhubCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployCharmhubCharm", reflect.TypeOf((*MockControllerCharmDeployer)(nil).DeployCharmhubCharm), arg0, arg1, arg2)
}

// DeployLocalCharm mocks base method.
func (m *MockControllerCharmDeployer) DeployLocalCharm(arg0 context.Context, arg1 string, arg2 base.Base) (string, *charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeployLocalCharm indicates an expected call of DeployLocalCharm.
func (mr *MockControllerCharmDeployerMockRecorder) DeployLocalCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployLocalCharm", reflect.TypeOf((*MockControllerCharmDeployer)(nil).DeployLocalCharm), arg0, arg1, arg2)
}

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPClient) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), arg0)
}

// MockLoggerFactory is a mock of LoggerFactory interface.
type MockLoggerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerFactoryMockRecorder
}

// MockLoggerFactoryMockRecorder is the mock recorder for MockLoggerFactory.
type MockLoggerFactoryMockRecorder struct {
	mock *MockLoggerFactory
}

// NewMockLoggerFactory creates a new mock instance.
func NewMockLoggerFactory(ctrl *gomock.Controller) *MockLoggerFactory {
	mock := &MockLoggerFactory{ctrl: ctrl}
	mock.recorder = &MockLoggerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerFactory) EXPECT() *MockLoggerFactoryMockRecorder {
	return m.recorder
}

// Child mocks base method.
func (m *MockLoggerFactory) Child(arg0 string) charmhub.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Child", arg0)
	ret0, _ := ret[0].(charmhub.Logger)
	return ret0
}

// Child indicates an expected call of Child.
func (mr *MockLoggerFactoryMockRecorder) Child(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Child", reflect.TypeOf((*MockLoggerFactory)(nil).Child), arg0)
}

// ChildWithLabels mocks base method.
func (m *MockLoggerFactory) ChildWithLabels(arg0 string, arg1 ...string) charmhub.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChildWithLabels", varargs...)
	ret0, _ := ret[0].(charmhub.Logger)
	return ret0
}

// ChildWithLabels indicates an expected call of ChildWithLabels.
func (mr *MockLoggerFactoryMockRecorder) ChildWithLabels(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildWithLabels", reflect.TypeOf((*MockLoggerFactory)(nil).ChildWithLabels), varargs...)
}

// Namespace mocks base method.
func (m *MockLoggerFactory) Namespace(arg0 string) services.LoggerFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace", arg0)
	ret0, _ := ret[0].(services.LoggerFactory)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockLoggerFactoryMockRecorder) Namespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockLoggerFactory)(nil).Namespace), arg0)
}

// MockCloudService is a mock of CloudService interface.
type MockCloudService struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceMockRecorder
}

// MockCloudServiceMockRecorder is the mock recorder for MockCloudService.
type MockCloudServiceMockRecorder struct {
	mock *MockCloudService
}

// NewMockCloudService creates a new mock instance.
func NewMockCloudService(ctrl *gomock.Controller) *MockCloudService {
	mock := &MockCloudService{ctrl: ctrl}
	mock.recorder = &MockCloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudService) EXPECT() *MockCloudServiceMockRecorder {
	return m.recorder
}

// Addresses mocks base method.
func (m *MockCloudService) Addresses() network.SpaceAddresses {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addresses")
	ret0, _ := ret[0].(network.SpaceAddresses)
	return ret0
}

// Addresses indicates an expected call of Addresses.
func (mr *MockCloudServiceMockRecorder) Addresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addresses", reflect.TypeOf((*MockCloudService)(nil).Addresses))
}

// MockCloudServiceGetter is a mock of CloudServiceGetter interface.
type MockCloudServiceGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCloudServiceGetterMockRecorder
}

// MockCloudServiceGetterMockRecorder is the mock recorder for MockCloudServiceGetter.
type MockCloudServiceGetterMockRecorder struct {
	mock *MockCloudServiceGetter
}

// NewMockCloudServiceGetter creates a new mock instance.
func NewMockCloudServiceGetter(ctrl *gomock.Controller) *MockCloudServiceGetter {
	mock := &MockCloudServiceGetter{ctrl: ctrl}
	mock.recorder = &MockCloudServiceGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudServiceGetter) EXPECT() *MockCloudServiceGetterMockRecorder {
	return m.recorder
}

// CloudService mocks base method.
func (m *MockCloudServiceGetter) CloudService(arg0 string) (CloudService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudService", arg0)
	ret0, _ := ret[0].(CloudService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudService indicates an expected call of CloudService.
func (mr *MockCloudServiceGetterMockRecorder) CloudService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudService", reflect.TypeOf((*MockCloudServiceGetter)(nil).CloudService), arg0)
}

// MockOperationApplier is a mock of OperationApplier interface.
type MockOperationApplier struct {
	ctrl     *gomock.Controller
	recorder *MockOperationApplierMockRecorder
}

// MockOperationApplierMockRecorder is the mock recorder for MockOperationApplier.
type MockOperationApplierMockRecorder struct {
	mock *MockOperationApplier
}

// NewMockOperationApplier creates a new mock instance.
func NewMockOperationApplier(ctrl *gomock.Controller) *MockOperationApplier {
	mock := &MockOperationApplier{ctrl: ctrl}
	mock.recorder = &MockOperationApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperationApplier) EXPECT() *MockOperationApplierMockRecorder {
	return m.recorder
}

// ApplyOperation mocks base method.
func (m *MockOperationApplier) ApplyOperation(arg0 *state.UpdateUnitOperation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyOperation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyOperation indicates an expected call of ApplyOperation.
func (mr *MockOperationApplierMockRecorder) ApplyOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyOperation", reflect.TypeOf((*MockOperationApplier)(nil).ApplyOperation), arg0)
}

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// AddPrincipal mocks base method.
func (m *MockMachine) AddPrincipal(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddPrincipal", arg0)
}

// AddPrincipal indicates an expected call of AddPrincipal.
func (mr *MockMachineMockRecorder) AddPrincipal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPrincipal", reflect.TypeOf((*MockMachine)(nil).AddPrincipal), arg0)
}

// Base mocks base method.
func (m *MockMachine) Base() state.Base {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Base")
	ret0, _ := ret[0].(state.Base)
	return ret0
}

// Base indicates an expected call of Base.
func (mr *MockMachineMockRecorder) Base() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Base", reflect.TypeOf((*MockMachine)(nil).Base))
}

// Clean mocks base method.
func (m *MockMachine) Clean() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clean")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Clean indicates an expected call of Clean.
func (mr *MockMachineMockRecorder) Clean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockMachine)(nil).Clean))
}

// ContainerType mocks base method.
func (m *MockMachine) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockMachineMockRecorder) ContainerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
}

// DocID mocks base method.
func (m *MockMachine) DocID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID")
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID.
func (mr *MockMachineMockRecorder) DocID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockMachine)(nil).DocID))
}

// FileSystems mocks base method.
func (m *MockMachine) FileSystems() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileSystems")
	ret0, _ := ret[0].([]string)
	return ret0
}

// FileSystems indicates an expected call of FileSystems.
func (mr *MockMachineMockRecorder) FileSystems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileSystems", reflect.TypeOf((*MockMachine)(nil).FileSystems))
}

// Id mocks base method.
func (m *MockMachine) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockMachineMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
}

// Jobs mocks base method.
func (m *MockMachine) Jobs() []state.MachineJob {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].([]state.MachineJob)
	return ret0
}

// Jobs indicates an expected call of Jobs.
func (mr *MockMachineMockRecorder) Jobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockMachine)(nil).Jobs))
}

// Life mocks base method.
func (m *MockMachine) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockMachineMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockMachine)(nil).Life))
}

// MachineTag mocks base method.
func (m *MockMachine) MachineTag() names.MachineTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineTag")
	ret0, _ := ret[0].(names.MachineTag)
	return ret0
}

// MachineTag indicates an expected call of MachineTag.
func (mr *MockMachineMockRecorder) MachineTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineTag", reflect.TypeOf((*MockMachine)(nil).MachineTag))
}

// PublicAddress mocks base method.
func (m *MockMachine) PublicAddress() (network.SpaceAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicAddress")
	ret0, _ := ret[0].(network.SpaceAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicAddress indicates an expected call of PublicAddress.
func (mr *MockMachineMockRecorder) PublicAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicAddress", reflect.TypeOf((*MockMachine)(nil).PublicAddress))
}

// MockMachineGetter is a mock of MachineGetter interface.
type MockMachineGetter struct {
	ctrl     *gomock.Controller
	recorder *MockMachineGetterMockRecorder
}

// MockMachineGetterMockRecorder is the mock recorder for MockMachineGetter.
type MockMachineGetterMockRecorder struct {
	mock *MockMachineGetter
}

// NewMockMachineGetter creates a new mock instance.
func NewMockMachineGetter(ctrl *gomock.Controller) *MockMachineGetter {
	mock := &MockMachineGetter{ctrl: ctrl}
	mock.recorder = &MockMachineGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineGetter) EXPECT() *MockMachineGetterMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockMachineGetter) Machine(arg0 string) (Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockMachineGetterMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockMachineGetter)(nil).Machine), arg0)
}

// MockStateBackend is a mock of StateBackend interface.
type MockStateBackend struct {
	ctrl     *gomock.Controller
	recorder *MockStateBackendMockRecorder
}

// MockStateBackendMockRecorder is the mock recorder for MockStateBackend.
type MockStateBackendMockRecorder struct {
	mock *MockStateBackend
}

// NewMockStateBackend creates a new mock instance.
func NewMockStateBackend(ctrl *gomock.Controller) *MockStateBackend {
	mock := &MockStateBackend{ctrl: ctrl}
	mock.recorder = &MockStateBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateBackend) EXPECT() *MockStateBackendMockRecorder {
	return m.recorder
}

// AddApplication mocks base method.
func (m *MockStateBackend) AddApplication(arg0 state.AddApplicationArgs, arg1 objectstore.ObjectStore) (Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddApplication", arg0, arg1)
	ret0, _ := ret[0].(Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddApplication indicates an expected call of AddApplication.
func (mr *MockStateBackendMockRecorder) AddApplication(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApplication", reflect.TypeOf((*MockStateBackend)(nil).AddApplication), arg0, arg1)
}

// Charm mocks base method.
func (m *MockStateBackend) Charm(arg0 string) (Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockStateBackendMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockStateBackend)(nil).Charm), arg0)
}

// Model mocks base method.
func (m *MockStateBackend) Model() (Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateBackendMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockStateBackend)(nil).Model))
}

// Unit mocks base method.
func (m *MockStateBackend) Unit(arg0 string) (Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockStateBackendMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockStateBackend)(nil).Unit), arg0)
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockApplication) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockApplicationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
}

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// Manifest mocks base method.
func (m *MockCharm) Manifest() *charm.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*charm.Manifest)
	return ret0
}

// Manifest indicates an expected call of Manifest.
func (mr *MockCharmMockRecorder) Manifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockCharm)(nil).Manifest))
}

// Meta mocks base method.
func (m *MockCharm) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockCharmMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharm)(nil).Meta))
}

// URL mocks base method.
func (m *MockCharm) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockCharmMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockCharm)(nil).URL))
}

// MockUnit is a mock of Unit interface.
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit.
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance.
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// AssignToMachineRef mocks base method.
func (m *MockUnit) AssignToMachineRef(arg0 state.MachineRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignToMachineRef", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignToMachineRef indicates an expected call of AssignToMachineRef.
func (mr *MockUnitMockRecorder) AssignToMachineRef(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignToMachineRef", reflect.TypeOf((*MockUnit)(nil).AssignToMachineRef), arg0)
}

// SetPassword mocks base method.
func (m *MockUnit) SetPassword(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPassword", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockUnitMockRecorder) SetPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockUnit)(nil).SetPassword), arg0)
}

// UnitTag mocks base method.
func (m *MockUnit) UnitTag() names.UnitTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitTag")
	ret0, _ := ret[0].(names.UnitTag)
	return ret0
}

// UnitTag indicates an expected call of UnitTag.
func (mr *MockUnitMockRecorder) UnitTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitTag", reflect.TypeOf((*MockUnit)(nil).UnitTag))
}

// UpdateOperation mocks base method.
func (m *MockUnit) UpdateOperation(arg0 state.UnitUpdateProperties) *state.UpdateUnitOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOperation", arg0)
	ret0, _ := ret[0].(*state.UpdateUnitOperation)
	return ret0
}

// UpdateOperation indicates an expected call of UpdateOperation.
func (mr *MockUnitMockRecorder) UpdateOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperation", reflect.TypeOf((*MockUnit)(nil).UpdateOperation), arg0)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModel)(nil).Config))
}

// MockCharmUploader is a mock of CharmUploader interface.
type MockCharmUploader struct {
	ctrl     *gomock.Controller
	recorder *MockCharmUploaderMockRecorder
}

// MockCharmUploaderMockRecorder is the mock recorder for MockCharmUploader.
type MockCharmUploaderMockRecorder struct {
	mock *MockCharmUploader
}

// NewMockCharmUploader creates a new mock instance.
func NewMockCharmUploader(ctrl *gomock.Controller) *MockCharmUploader {
	mock := &MockCharmUploader{ctrl: ctrl}
	mock.recorder = &MockCharmUploaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmUploader) EXPECT() *MockCharmUploaderMockRecorder {
	return m.recorder
}

// ModelUUID mocks base method.
func (m *MockCharmUploader) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockCharmUploaderMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockCharmUploader)(nil).ModelUUID))
}

// PrepareCharmUpload mocks base method.
func (m *MockCharmUploader) PrepareCharmUpload(arg0 string) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockCharmUploaderMockRecorder) PrepareCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockCharmUploader)(nil).PrepareCharmUpload), arg0)
}

// PrepareLocalCharmUpload mocks base method.
func (m *MockCharmUploader) PrepareLocalCharmUpload(arg0 string) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareLocalCharmUpload", arg0)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareLocalCharmUpload indicates an expected call of PrepareLocalCharmUpload.
func (mr *MockCharmUploaderMockRecorder) PrepareLocalCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareLocalCharmUpload", reflect.TypeOf((*MockCharmUploader)(nil).PrepareLocalCharmUpload), arg0)
}

// UpdateUploadedCharm mocks base method.
func (m *MockCharmUploader) UpdateUploadedCharm(arg0 state.CharmInfo) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockCharmUploaderMockRecorder) UpdateUploadedCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockCharmUploader)(nil).UpdateUploadedCharm), arg0)
}
