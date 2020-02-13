// Code generated by MockGen. DO NOT EDIT.
// Source: startup.go

// Package mock_startup is a generated GoMock package.
package mock_startup

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"

	api "github.com/openshift/openshift-azure/pkg/api"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// WriteFiles mocks base method
func (m *MockInterface) WriteFiles(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFiles", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFiles indicates an expected call of WriteFiles
func (mr *MockInterfaceMockRecorder) WriteFiles(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFiles", reflect.TypeOf((*MockInterface)(nil).WriteFiles), ctx)
}

// Hash mocks base method
func (m *MockInterface) Hash(role api.AgentPoolProfileRole) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", role)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (mr *MockInterfaceMockRecorder) Hash(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockInterface)(nil).Hash), role)
}

// GetWorkerCs mocks base method
func (m *MockInterface) GetWorkerCs() *api.OpenShiftManagedCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkerCs")
	ret0, _ := ret[0].(*api.OpenShiftManagedCluster)
	return ret0
}

// GetWorkerCs indicates an expected call of GetWorkerCs
func (mr *MockInterfaceMockRecorder) GetWorkerCs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkerCs", reflect.TypeOf((*MockInterface)(nil).GetWorkerCs))
}

// WriteSearchDomain mocks base method
func (m *MockInterface) WriteSearchDomain(ctx context.Context, log *logrus.Entry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteSearchDomain", ctx, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteSearchDomain indicates an expected call of WriteSearchDomain
func (mr *MockInterfaceMockRecorder) WriteSearchDomain(ctx, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSearchDomain", reflect.TypeOf((*MockInterface)(nil).WriteSearchDomain), ctx, log)
}
