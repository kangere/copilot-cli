// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/runner/jobrunner/job_runner.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cloudformation "github.com/aws/copilot-cli/internal/pkg/aws/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

// MockjobExecutor is a mock of jobExecutor interface.
type MockjobExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockjobExecutorMockRecorder
}

// MockjobExecutorMockRecorder is the mock recorder for MockjobExecutor.
type MockjobExecutorMockRecorder struct {
	mock *MockjobExecutor
}

// NewMockjobExecutor creates a new mock instance.
func NewMockjobExecutor(ctrl *gomock.Controller) *MockjobExecutor {
	mock := &MockjobExecutor{ctrl: ctrl}
	mock.recorder = &MockjobExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockjobExecutor) EXPECT() *MockjobExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockjobExecutor) Execute(stateMachineARN string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", stateMachineARN)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockjobExecutorMockRecorder) Execute(stateMachineARN interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockjobExecutor)(nil).Execute), stateMachineARN)
}

// MockStackRetriever is a mock of StackRetriever interface.
type MockStackRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockStackRetrieverMockRecorder
}

// MockStackRetrieverMockRecorder is the mock recorder for MockStackRetriever.
type MockStackRetrieverMockRecorder struct {
	mock *MockStackRetriever
}

// NewMockStackRetriever creates a new mock instance.
func NewMockStackRetriever(ctrl *gomock.Controller) *MockStackRetriever {
	mock := &MockStackRetriever{ctrl: ctrl}
	mock.recorder = &MockStackRetrieverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStackRetriever) EXPECT() *MockStackRetrieverMockRecorder {
	return m.recorder
}

// StackResources mocks base method.
func (m *MockStackRetriever) StackResources(name string) ([]*cloudformation.StackResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackResources", name)
	ret0, _ := ret[0].([]*cloudformation.StackResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackResources indicates an expected call of StackResources.
func (mr *MockStackRetrieverMockRecorder) StackResources(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackResources", reflect.TypeOf((*MockStackRetriever)(nil).StackResources), name)
}
