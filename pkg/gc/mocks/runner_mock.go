// Code generated by MockGen. DO NOT EDIT.
// Source: ../../pkg/gc/task.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// RunGC mocks base method.
func (m *MockRunner) RunGC() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunGC")
	ret0, _ := ret[0].(error)
	return ret0
}

// RunGC indicates an expected call of RunGC.
func (mr *MockRunnerMockRecorder) RunGC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunGC", reflect.TypeOf((*MockRunner)(nil).RunGC))
}
