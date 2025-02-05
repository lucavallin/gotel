// Code generated by mockery v2.50.0. DO NOT EDIT.

package trace

import (
	mock "github.com/stretchr/testify/mock"
	trace "go.opentelemetry.io/otel/trace"
)

// MockSpanOption is an autogenerated mock type for the SpanOption type
type MockSpanOption struct {
	mock.Mock
}

type MockSpanOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpanOption) EXPECT() *MockSpanOption_Expecter {
	return &MockSpanOption_Expecter{mock: &_m.Mock}
}

// applySpanEnd provides a mock function with given fields: _a0
func (_m *MockSpanOption) applySpanEnd(_a0 trace.SpanConfig) trace.SpanConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applySpanEnd")
	}

	var r0 trace.SpanConfig
	if rf, ok := ret.Get(0).(func(trace.SpanConfig) trace.SpanConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(trace.SpanConfig)
	}

	return r0
}

// MockSpanOption_applySpanEnd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applySpanEnd'
type MockSpanOption_applySpanEnd_Call struct {
	*mock.Call
}

// applySpanEnd is a helper method to define mock.On call
//   - _a0 trace.SpanConfig
func (_e *MockSpanOption_Expecter) applySpanEnd(_a0 interface{}) *MockSpanOption_applySpanEnd_Call {
	return &MockSpanOption_applySpanEnd_Call{Call: _e.mock.On("applySpanEnd", _a0)}
}

func (_c *MockSpanOption_applySpanEnd_Call) Run(run func(_a0 trace.SpanConfig)) *MockSpanOption_applySpanEnd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(trace.SpanConfig))
	})
	return _c
}

func (_c *MockSpanOption_applySpanEnd_Call) Return(_a0 trace.SpanConfig) *MockSpanOption_applySpanEnd_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpanOption_applySpanEnd_Call) RunAndReturn(run func(trace.SpanConfig) trace.SpanConfig) *MockSpanOption_applySpanEnd_Call {
	_c.Call.Return(run)
	return _c
}

// applySpanStart provides a mock function with given fields: _a0
func (_m *MockSpanOption) applySpanStart(_a0 trace.SpanConfig) trace.SpanConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applySpanStart")
	}

	var r0 trace.SpanConfig
	if rf, ok := ret.Get(0).(func(trace.SpanConfig) trace.SpanConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(trace.SpanConfig)
	}

	return r0
}

// MockSpanOption_applySpanStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applySpanStart'
type MockSpanOption_applySpanStart_Call struct {
	*mock.Call
}

// applySpanStart is a helper method to define mock.On call
//   - _a0 trace.SpanConfig
func (_e *MockSpanOption_Expecter) applySpanStart(_a0 interface{}) *MockSpanOption_applySpanStart_Call {
	return &MockSpanOption_applySpanStart_Call{Call: _e.mock.On("applySpanStart", _a0)}
}

func (_c *MockSpanOption_applySpanStart_Call) Run(run func(_a0 trace.SpanConfig)) *MockSpanOption_applySpanStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(trace.SpanConfig))
	})
	return _c
}

func (_c *MockSpanOption_applySpanStart_Call) Return(_a0 trace.SpanConfig) *MockSpanOption_applySpanStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpanOption_applySpanStart_Call) RunAndReturn(run func(trace.SpanConfig) trace.SpanConfig) *MockSpanOption_applySpanStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpanOption creates a new instance of MockSpanOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpanOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpanOption {
	mock := &MockSpanOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
