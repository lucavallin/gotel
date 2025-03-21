// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockTracerProvider is an autogenerated mock type for the TracerProvider type
type MockTracerProvider struct {
	mock.Mock
}

type MockTracerProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTracerProvider) EXPECT() *MockTracerProvider_Expecter {
	return &MockTracerProvider_Expecter{mock: &_m.Mock}
}

// tracerProvider provides a mock function with no fields
func (_m *MockTracerProvider) tracerProvider() {
	_m.Called()
}

// MockTracerProvider_tracerProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'tracerProvider'
type MockTracerProvider_tracerProvider_Call struct {
	*mock.Call
}

// tracerProvider is a helper method to define mock.On call
func (_e *MockTracerProvider_Expecter) tracerProvider() *MockTracerProvider_tracerProvider_Call {
	return &MockTracerProvider_tracerProvider_Call{Call: _e.mock.On("tracerProvider")}
}

func (_c *MockTracerProvider_tracerProvider_Call) Run(run func()) *MockTracerProvider_tracerProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTracerProvider_tracerProvider_Call) Return() *MockTracerProvider_tracerProvider_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTracerProvider_tracerProvider_Call) RunAndReturn(run func()) *MockTracerProvider_tracerProvider_Call {
	_c.Run(run)
	return _c
}

// NewMockTracerProvider creates a new instance of MockTracerProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTracerProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTracerProvider {
	mock := &MockTracerProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
