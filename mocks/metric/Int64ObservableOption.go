// Code generated by mockery v2.50.0. DO NOT EDIT.

package metric

import (
	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockInt64ObservableOption is an autogenerated mock type for the Int64ObservableOption type
type MockInt64ObservableOption struct {
	mock.Mock
}

type MockInt64ObservableOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInt64ObservableOption) EXPECT() *MockInt64ObservableOption_Expecter {
	return &MockInt64ObservableOption_Expecter{mock: &_m.Mock}
}

// applyInt64ObservableCounter provides a mock function with given fields: _a0
func (_m *MockInt64ObservableOption) applyInt64ObservableCounter(_a0 metric.Int64ObservableCounterConfig) metric.Int64ObservableCounterConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyInt64ObservableCounter")
	}

	var r0 metric.Int64ObservableCounterConfig
	if rf, ok := ret.Get(0).(func(metric.Int64ObservableCounterConfig) metric.Int64ObservableCounterConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.Int64ObservableCounterConfig)
	}

	return r0
}

// MockInt64ObservableOption_applyInt64ObservableCounter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyInt64ObservableCounter'
type MockInt64ObservableOption_applyInt64ObservableCounter_Call struct {
	*mock.Call
}

// applyInt64ObservableCounter is a helper method to define mock.On call
//   - _a0 metric.Int64ObservableCounterConfig
func (_e *MockInt64ObservableOption_Expecter) applyInt64ObservableCounter(_a0 interface{}) *MockInt64ObservableOption_applyInt64ObservableCounter_Call {
	return &MockInt64ObservableOption_applyInt64ObservableCounter_Call{Call: _e.mock.On("applyInt64ObservableCounter", _a0)}
}

func (_c *MockInt64ObservableOption_applyInt64ObservableCounter_Call) Run(run func(_a0 metric.Int64ObservableCounterConfig)) *MockInt64ObservableOption_applyInt64ObservableCounter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Int64ObservableCounterConfig))
	})
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableCounter_Call) Return(_a0 metric.Int64ObservableCounterConfig) *MockInt64ObservableOption_applyInt64ObservableCounter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableCounter_Call) RunAndReturn(run func(metric.Int64ObservableCounterConfig) metric.Int64ObservableCounterConfig) *MockInt64ObservableOption_applyInt64ObservableCounter_Call {
	_c.Call.Return(run)
	return _c
}

// applyInt64ObservableGauge provides a mock function with given fields: _a0
func (_m *MockInt64ObservableOption) applyInt64ObservableGauge(_a0 metric.Int64ObservableGaugeConfig) metric.Int64ObservableGaugeConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyInt64ObservableGauge")
	}

	var r0 metric.Int64ObservableGaugeConfig
	if rf, ok := ret.Get(0).(func(metric.Int64ObservableGaugeConfig) metric.Int64ObservableGaugeConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.Int64ObservableGaugeConfig)
	}

	return r0
}

// MockInt64ObservableOption_applyInt64ObservableGauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyInt64ObservableGauge'
type MockInt64ObservableOption_applyInt64ObservableGauge_Call struct {
	*mock.Call
}

// applyInt64ObservableGauge is a helper method to define mock.On call
//   - _a0 metric.Int64ObservableGaugeConfig
func (_e *MockInt64ObservableOption_Expecter) applyInt64ObservableGauge(_a0 interface{}) *MockInt64ObservableOption_applyInt64ObservableGauge_Call {
	return &MockInt64ObservableOption_applyInt64ObservableGauge_Call{Call: _e.mock.On("applyInt64ObservableGauge", _a0)}
}

func (_c *MockInt64ObservableOption_applyInt64ObservableGauge_Call) Run(run func(_a0 metric.Int64ObservableGaugeConfig)) *MockInt64ObservableOption_applyInt64ObservableGauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Int64ObservableGaugeConfig))
	})
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableGauge_Call) Return(_a0 metric.Int64ObservableGaugeConfig) *MockInt64ObservableOption_applyInt64ObservableGauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableGauge_Call) RunAndReturn(run func(metric.Int64ObservableGaugeConfig) metric.Int64ObservableGaugeConfig) *MockInt64ObservableOption_applyInt64ObservableGauge_Call {
	_c.Call.Return(run)
	return _c
}

// applyInt64ObservableUpDownCounter provides a mock function with given fields: _a0
func (_m *MockInt64ObservableOption) applyInt64ObservableUpDownCounter(_a0 metric.Int64ObservableUpDownCounterConfig) metric.Int64ObservableUpDownCounterConfig {
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

// MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyInt64ObservableUpDownCounter'
type MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call struct {
	*mock.Call
}

// applyInt64ObservableUpDownCounter is a helper method to define mock.On call
//   - _a0 metric.Int64ObservableUpDownCounterConfig
func (_e *MockInt64ObservableOption_Expecter) applyInt64ObservableUpDownCounter(_a0 interface{}) *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call {
	return &MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call{Call: _e.mock.On("applyInt64ObservableUpDownCounter", _a0)}
}

func (_c *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call) Run(run func(_a0 metric.Int64ObservableUpDownCounterConfig)) *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Int64ObservableUpDownCounterConfig))
	})
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call) Return(_a0 metric.Int64ObservableUpDownCounterConfig) *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call) RunAndReturn(run func(metric.Int64ObservableUpDownCounterConfig) metric.Int64ObservableUpDownCounterConfig) *MockInt64ObservableOption_applyInt64ObservableUpDownCounter_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInt64ObservableOption creates a new instance of MockInt64ObservableOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInt64ObservableOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInt64ObservableOption {
	mock := &MockInt64ObservableOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
