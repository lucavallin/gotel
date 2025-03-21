// Code generated by mockery v2.50.0. DO NOT EDIT.

package trace

import (
	mock "github.com/stretchr/testify/mock"
	trace "go.opentelemetry.io/otel/trace"
)

// MockspanOptionFunc is an autogenerated mock type for the spanOptionFunc type
type MockspanOptionFunc struct {
	mock.Mock
}

type MockspanOptionFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockspanOptionFunc) EXPECT() *MockspanOptionFunc_Expecter {
	return &MockspanOptionFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *MockspanOptionFunc) Execute(_a0 trace.SpanConfig) trace.SpanConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 trace.SpanConfig
	if rf, ok := ret.Get(0).(func(trace.SpanConfig) trace.SpanConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(trace.SpanConfig)
	}

	return r0
}

// MockspanOptionFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockspanOptionFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 trace.SpanConfig
func (_e *MockspanOptionFunc_Expecter) Execute(_a0 interface{}) *MockspanOptionFunc_Execute_Call {
	return &MockspanOptionFunc_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *MockspanOptionFunc_Execute_Call) Run(run func(_a0 trace.SpanConfig)) *MockspanOptionFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(trace.SpanConfig))
	})
	return _c
}

func (_c *MockspanOptionFunc_Execute_Call) Return(_a0 trace.SpanConfig) *MockspanOptionFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockspanOptionFunc_Execute_Call) RunAndReturn(run func(trace.SpanConfig) trace.SpanConfig) *MockspanOptionFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockspanOptionFunc creates a new instance of MockspanOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockspanOptionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockspanOptionFunc {
	mock := &MockspanOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
