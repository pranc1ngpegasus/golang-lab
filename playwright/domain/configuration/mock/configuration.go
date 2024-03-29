// Code generated by MockGen. DO NOT EDIT.
// Source: configuration.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	configuration "github.com/Pranc1ngPegasus/golang-lab/playwright/domain/configuration"
	gomock "github.com/golang/mock/gomock"
)

// MockConfiguration is a mock of Configuration interface.
type MockConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationMockRecorder
}

// MockConfigurationMockRecorder is the mock recorder for MockConfiguration.
type MockConfigurationMockRecorder struct {
	mock *MockConfiguration
}

// NewMockConfiguration creates a new mock instance.
func NewMockConfiguration(ctrl *gomock.Controller) *MockConfiguration {
	mock := &MockConfiguration{ctrl: ctrl}
	mock.recorder = &MockConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfiguration) EXPECT() *MockConfigurationMockRecorder {
	return m.recorder
}

// Common mocks base method.
func (m *MockConfiguration) Common() *configuration.Common {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Common")
	ret0, _ := ret[0].(*configuration.Common)
	return ret0
}

// Common indicates an expected call of Common.
func (mr *MockConfigurationMockRecorder) Common() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Common", reflect.TypeOf((*MockConfiguration)(nil).Common))
}
