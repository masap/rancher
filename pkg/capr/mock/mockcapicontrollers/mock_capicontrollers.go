// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/rancher/pkg/generated/controllers/cluster.x-k8s.io/v1beta1 (interfaces: MachineClient,MachineCache,ClusterClient,ClusterCache)

// Package mockcapicontrollers is a generated GoMock package.
package mockcapicontrollers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/rancher/rancher/pkg/generated/controllers/cluster.x-k8s.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockMachineClient is a mock of MachineClient interface.
type MockMachineClient struct {
	ctrl     *gomock.Controller
	recorder *MockMachineClientMockRecorder
}

// MockMachineClientMockRecorder is the mock recorder for MockMachineClient.
type MockMachineClientMockRecorder struct {
	mock *MockMachineClient
}

// NewMockMachineClient creates a new mock instance.
func NewMockMachineClient(ctrl *gomock.Controller) *MockMachineClient {
	mock := &MockMachineClient{ctrl: ctrl}
	mock.recorder = &MockMachineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineClient) EXPECT() *MockMachineClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMachineClient) Create(arg0 *v1beta10.Machine) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMachineClientMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMachineClient)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockMachineClient) Delete(arg0, arg1 string, arg2 *v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMachineClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMachineClient)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMachineClient) Get(arg0, arg1 string, arg2 v1.GetOptions) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMachineClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMachineClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMachineClient) List(arg0 string, arg1 v1.ListOptions) (*v1beta10.MachineList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta10.MachineList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMachineClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMachineClient)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockMachineClient) Patch(arg0, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 ...string) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockMachineClientMockRecorder) Patch(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockMachineClient)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockMachineClient) Update(arg0 *v1beta10.Machine) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMachineClientMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMachineClient)(nil).Update), arg0)
}

// UpdateStatus mocks base method.
func (m *MockMachineClient) UpdateStatus(arg0 *v1beta10.Machine) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockMachineClientMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockMachineClient)(nil).UpdateStatus), arg0)
}

// Watch mocks base method.
func (m *MockMachineClient) Watch(arg0 string, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockMachineClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMachineClient)(nil).Watch), arg0, arg1)
}

// MockMachineCache is a mock of MachineCache interface.
type MockMachineCache struct {
	ctrl     *gomock.Controller
	recorder *MockMachineCacheMockRecorder
}

// MockMachineCacheMockRecorder is the mock recorder for MockMachineCache.
type MockMachineCacheMockRecorder struct {
	mock *MockMachineCache
}

// NewMockMachineCache creates a new mock instance.
func NewMockMachineCache(ctrl *gomock.Controller) *MockMachineCache {
	mock := &MockMachineCache{ctrl: ctrl}
	mock.recorder = &MockMachineCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineCache) EXPECT() *MockMachineCacheMockRecorder {
	return m.recorder
}

// AddIndexer mocks base method.
func (m *MockMachineCache) AddIndexer(arg0 string, arg1 v1beta1.MachineIndexer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIndexer", arg0, arg1)
}

// AddIndexer indicates an expected call of AddIndexer.
func (mr *MockMachineCacheMockRecorder) AddIndexer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexer", reflect.TypeOf((*MockMachineCache)(nil).AddIndexer), arg0, arg1)
}

// Get mocks base method.
func (m *MockMachineCache) Get(arg0, arg1 string) (*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMachineCacheMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMachineCache)(nil).Get), arg0, arg1)
}

// GetByIndex mocks base method.
func (m *MockMachineCache) GetByIndex(arg0, arg1 string) ([]*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0, arg1)
	ret0, _ := ret[0].([]*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockMachineCacheMockRecorder) GetByIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockMachineCache)(nil).GetByIndex), arg0, arg1)
}

// List mocks base method.
func (m *MockMachineCache) List(arg0 string, arg1 labels.Selector) ([]*v1beta10.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*v1beta10.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMachineCacheMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMachineCache)(nil).List), arg0, arg1)
}

// MockClusterClient is a mock of ClusterClient interface.
type MockClusterClient struct {
	ctrl     *gomock.Controller
	recorder *MockClusterClientMockRecorder
}

// MockClusterClientMockRecorder is the mock recorder for MockClusterClient.
type MockClusterClientMockRecorder struct {
	mock *MockClusterClient
}

// NewMockClusterClient creates a new mock instance.
func NewMockClusterClient(ctrl *gomock.Controller) *MockClusterClient {
	mock := &MockClusterClient{ctrl: ctrl}
	mock.recorder = &MockClusterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterClient) EXPECT() *MockClusterClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClusterClient) Create(arg0 *v1beta10.Cluster) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockClusterClientMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterClient)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockClusterClient) Delete(arg0, arg1 string, arg2 *v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClusterClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterClient)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockClusterClient) Get(arg0, arg1 string, arg2 v1.GetOptions) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterClient)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockClusterClient) List(arg0 string, arg1 v1.ListOptions) (*v1beta10.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1beta10.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterClient)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockClusterClient) Patch(arg0, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 ...string) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockClusterClientMockRecorder) Patch(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockClusterClient)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockClusterClient) Update(arg0 *v1beta10.Cluster) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockClusterClientMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterClient)(nil).Update), arg0)
}

// UpdateStatus mocks base method.
func (m *MockClusterClient) UpdateStatus(arg0 *v1beta10.Cluster) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockClusterClientMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClusterClient)(nil).UpdateStatus), arg0)
}

// Watch mocks base method.
func (m *MockClusterClient) Watch(arg0 string, arg1 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockClusterClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterClient)(nil).Watch), arg0, arg1)
}

// MockClusterCache is a mock of ClusterCache interface.
type MockClusterCache struct {
	ctrl     *gomock.Controller
	recorder *MockClusterCacheMockRecorder
}

// MockClusterCacheMockRecorder is the mock recorder for MockClusterCache.
type MockClusterCacheMockRecorder struct {
	mock *MockClusterCache
}

// NewMockClusterCache creates a new mock instance.
func NewMockClusterCache(ctrl *gomock.Controller) *MockClusterCache {
	mock := &MockClusterCache{ctrl: ctrl}
	mock.recorder = &MockClusterCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterCache) EXPECT() *MockClusterCacheMockRecorder {
	return m.recorder
}

// AddIndexer mocks base method.
func (m *MockClusterCache) AddIndexer(arg0 string, arg1 v1beta1.ClusterIndexer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIndexer", arg0, arg1)
}

// AddIndexer indicates an expected call of AddIndexer.
func (mr *MockClusterCacheMockRecorder) AddIndexer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexer", reflect.TypeOf((*MockClusterCache)(nil).AddIndexer), arg0, arg1)
}

// Get mocks base method.
func (m *MockClusterCache) Get(arg0, arg1 string) (*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockClusterCacheMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterCache)(nil).Get), arg0, arg1)
}

// GetByIndex mocks base method.
func (m *MockClusterCache) GetByIndex(arg0, arg1 string) ([]*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", arg0, arg1)
	ret0, _ := ret[0].([]*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockClusterCacheMockRecorder) GetByIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockClusterCache)(nil).GetByIndex), arg0, arg1)
}

// List mocks base method.
func (m *MockClusterCache) List(arg0 string, arg1 labels.Selector) ([]*v1beta10.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*v1beta10.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockClusterCacheMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterCache)(nil).List), arg0, arg1)
}