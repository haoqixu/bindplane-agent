// Code generated by mockery v2.37.1. DO NOT EDIT.

package chronicleexporter

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	plog "go.opentelemetry.io/collector/pdata/plog"
)

// MockMarshaler is an autogenerated mock type for the logMarshaler type
type MockMarshaler struct {
	mock.Mock
}

type MockMarshaler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMarshaler) EXPECT() *MockMarshaler_Expecter {
	return &MockMarshaler_Expecter{mock: &_m.Mock}
}

// MarshalRawLogs provides a mock function with given fields: ctx, ld
func (_m *MockMarshaler) MarshalRawLogs(ctx context.Context, ld plog.Logs) ([]payload, error) {
	ret := _m.Called(ctx, ld)

	var r0 []payload
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, plog.Logs) ([]payload, error)); ok {
		return rf(ctx, ld)
	}
	if rf, ok := ret.Get(0).(func(context.Context, plog.Logs) []payload); ok {
		r0 = rf(ctx, ld)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payload)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, plog.Logs) error); ok {
		r1 = rf(ctx, ld)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMarshaler_MarshalRawLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalRawLogs'
type MockMarshaler_MarshalRawLogs_Call struct {
	*mock.Call
}

// MarshalRawLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - ld plog.Logs
func (_e *MockMarshaler_Expecter) MarshalRawLogs(ctx interface{}, ld interface{}) *MockMarshaler_MarshalRawLogs_Call {
	return &MockMarshaler_MarshalRawLogs_Call{Call: _e.mock.On("MarshalRawLogs", ctx, ld)}
}

func (_c *MockMarshaler_MarshalRawLogs_Call) Run(run func(ctx context.Context, ld plog.Logs)) *MockMarshaler_MarshalRawLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(plog.Logs))
	})
	return _c
}

func (_c *MockMarshaler_MarshalRawLogs_Call) Return(_a0 []payload, _a1 error) *MockMarshaler_MarshalRawLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMarshaler_MarshalRawLogs_Call) RunAndReturn(run func(context.Context, plog.Logs) ([]payload, error)) *MockMarshaler_MarshalRawLogs_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMarshaler creates a new instance of MockMarshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMarshaler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMarshaler {
	mock := &MockMarshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
