// Code generated by mockery v2.50.0. DO NOT EDIT.

package metric

import (
	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockFloat64CounterOption is an autogenerated mock type for the Float64CounterOption type
type MockFloat64CounterOption struct {
	mock.Mock
}

type MockFloat64CounterOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64CounterOption) EXPECT() *MockFloat64CounterOption_Expecter {
	return &MockFloat64CounterOption_Expecter{mock: &_m.Mock}
}

// applyFloat64Counter provides a mock function with given fields: _a0
func (_m *MockFloat64CounterOption) applyFloat64Counter(_a0 metric.Float64CounterConfig) metric.Float64CounterConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyFloat64Counter")
	}

	var r0 metric.Float64CounterConfig
	if rf, ok := ret.Get(0).(func(metric.Float64CounterConfig) metric.Float64CounterConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.Float64CounterConfig)
	}

	return r0
}

// MockFloat64CounterOption_applyFloat64Counter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyFloat64Counter'
type MockFloat64CounterOption_applyFloat64Counter_Call struct {
	*mock.Call
}

// applyFloat64Counter is a helper method to define mock.On call
//   - _a0 metric.Float64CounterConfig
func (_e *MockFloat64CounterOption_Expecter) applyFloat64Counter(_a0 interface{}) *MockFloat64CounterOption_applyFloat64Counter_Call {
	return &MockFloat64CounterOption_applyFloat64Counter_Call{Call: _e.mock.On("applyFloat64Counter", _a0)}
}

func (_c *MockFloat64CounterOption_applyFloat64Counter_Call) Run(run func(_a0 metric.Float64CounterConfig)) *MockFloat64CounterOption_applyFloat64Counter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Float64CounterConfig))
	})
	return _c
}

func (_c *MockFloat64CounterOption_applyFloat64Counter_Call) Return(_a0 metric.Float64CounterConfig) *MockFloat64CounterOption_applyFloat64Counter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFloat64CounterOption_applyFloat64Counter_Call) RunAndReturn(run func(metric.Float64CounterConfig) metric.Float64CounterConfig) *MockFloat64CounterOption_applyFloat64Counter_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFloat64CounterOption creates a new instance of MockFloat64CounterOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64CounterOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64CounterOption {
	mock := &MockFloat64CounterOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
