// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockInt64Gauge is an autogenerated mock type for the Int64Gauge type
type MockInt64Gauge struct {
	mock.Mock
}

type MockInt64Gauge_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInt64Gauge) EXPECT() *MockInt64Gauge_Expecter {
	return &MockInt64Gauge_Expecter{mock: &_m.Mock}
}

// int64Gauge provides a mock function with no fields
func (_m *MockInt64Gauge) int64Gauge() {
	_m.Called()
}

// MockInt64Gauge_int64Gauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'int64Gauge'
type MockInt64Gauge_int64Gauge_Call struct {
	*mock.Call
}

// int64Gauge is a helper method to define mock.On call
func (_e *MockInt64Gauge_Expecter) int64Gauge() *MockInt64Gauge_int64Gauge_Call {
	return &MockInt64Gauge_int64Gauge_Call{Call: _e.mock.On("int64Gauge")}
}

func (_c *MockInt64Gauge_int64Gauge_Call) Run(run func()) *MockInt64Gauge_int64Gauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInt64Gauge_int64Gauge_Call) Return() *MockInt64Gauge_int64Gauge_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInt64Gauge_int64Gauge_Call) RunAndReturn(run func()) *MockInt64Gauge_int64Gauge_Call {
	_c.Run(run)
	return _c
}

// NewMockInt64Gauge creates a new instance of MockInt64Gauge. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInt64Gauge(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInt64Gauge {
	mock := &MockInt64Gauge{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
