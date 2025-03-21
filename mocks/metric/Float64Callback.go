// Code generated by mockery v2.50.0. DO NOT EDIT.

package metric

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockFloat64Callback is an autogenerated mock type for the Float64Callback type
type MockFloat64Callback struct {
	mock.Mock
}

type MockFloat64Callback_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64Callback) EXPECT() *MockFloat64Callback_Expecter {
	return &MockFloat64Callback_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockFloat64Callback) Execute(_a0 context.Context, _a1 metric.Float64Observer) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, metric.Float64Observer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFloat64Callback_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockFloat64Callback_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 metric.Float64Observer
func (_e *MockFloat64Callback_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockFloat64Callback_Execute_Call {
	return &MockFloat64Callback_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockFloat64Callback_Execute_Call) Run(run func(_a0 context.Context, _a1 metric.Float64Observer)) *MockFloat64Callback_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(metric.Float64Observer))
	})
	return _c
}

func (_c *MockFloat64Callback_Execute_Call) Return(_a0 error) *MockFloat64Callback_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFloat64Callback_Execute_Call) RunAndReturn(run func(context.Context, metric.Float64Observer) error) *MockFloat64Callback_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFloat64Callback creates a new instance of MockFloat64Callback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64Callback(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64Callback {
	mock := &MockFloat64Callback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
