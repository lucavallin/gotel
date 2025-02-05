// Code generated by mockery v2.50.0. DO NOT EDIT.

package metric

import (
	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockInt64ObservableUpDownCounterOption is an autogenerated mock type for the Int64ObservableUpDownCounterOption type
type MockInt64ObservableUpDownCounterOption struct {
	mock.Mock
}

type MockInt64ObservableUpDownCounterOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInt64ObservableUpDownCounterOption) EXPECT() *MockInt64ObservableUpDownCounterOption_Expecter {
	return &MockInt64ObservableUpDownCounterOption_Expecter{mock: &_m.Mock}
}

// applyInt64ObservableUpDownCounter provides a mock function with given fields: _a0
func (_m *MockInt64ObservableUpDownCounterOption) applyInt64ObservableUpDownCounter(_a0 metric.Int64ObservableUpDownCounterConfig) metric.Int64ObservableUpDownCounterConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyInt64ObservableUpDownCounter")
	}

	var r0 metric.Int64ObservableUpDownCounterConfig
	if rf, ok := ret.Get(0).(func(metric.Int64ObservableUpDownCounterConfig) metric.Int64ObservableUpDownCounterConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.Int64ObservableUpDownCounterConfig)
	}

	return r0
}

// MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyInt64ObservableUpDownCounter'
type MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call struct {
	*mock.Call
}

// applyInt64ObservableUpDownCounter is a helper method to define mock.On call
//   - _a0 metric.Int64ObservableUpDownCounterConfig
func (_e *MockInt64ObservableUpDownCounterOption_Expecter) applyInt64ObservableUpDownCounter(_a0 interface{}) *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call {
	return &MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call{Call: _e.mock.On("applyInt64ObservableUpDownCounter", _a0)}
}

func (_c *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call) Run(run func(_a0 metric.Int64ObservableUpDownCounterConfig)) *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Int64ObservableUpDownCounterConfig))
	})
	return _c
}

func (_c *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call) Return(_a0 metric.Int64ObservableUpDownCounterConfig) *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call) RunAndReturn(run func(metric.Int64ObservableUpDownCounterConfig) metric.Int64ObservableUpDownCounterConfig) *MockInt64ObservableUpDownCounterOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInt64ObservableUpDownCounterOption creates a new instance of MockInt64ObservableUpDownCounterOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInt64ObservableUpDownCounterOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInt64ObservableUpDownCounterOption {
	mock := &MockInt64ObservableUpDownCounterOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
