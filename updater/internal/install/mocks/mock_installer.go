// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	rollback "github.com/observiq/observiq-otel-collector/updater/internal/rollback"
	mock "github.com/stretchr/testify/mock"
)

// MockInstaller is an autogenerated mock type for the Installer type
type MockInstaller struct {
	mock.Mock
}

// Install provides a mock function with given fields: _a0
func (_m *MockInstaller) Install(_a0 rollback.Rollbacker) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(rollback.Rollbacker) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockInstaller interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockInstaller creates a new instance of MockInstaller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockInstaller(t mockConstructorTestingTNewMockInstaller) *MockInstaller {
	mock := &MockInstaller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
