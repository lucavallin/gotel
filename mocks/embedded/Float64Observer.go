// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockFloat64Observer is an autogenerated mock type for the Float64Observer type
type MockFloat64Observer struct {
	mock.Mock
}

type MockFloat64Observer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64Observer) EXPECT() *MockFloat64Observer_Expecter {
	return &MockFloat64Observer_Expecter{mock: &_m.Mock}
}

// float64Observer provides a mock function with no fields
func (_m *MockFloat64Observer) float64Observer() {
	_m.Called()
}

// MockFloat64Observer_float64Observer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'float64Observer'
type MockFloat64Observer_float64Observer_Call struct {
	*mock.Call
}

// float64Observer is a helper method to define mock.On call
func (_e *MockFloat64Observer_Expecter) float64Observer() *MockFloat64Observer_float64Observer_Call {
	return &MockFloat64Observer_float64Observer_Call{Call: _e.mock.On("float64Observer")}
}

func (_c *MockFloat64Observer_float64Observer_Call) Run(run func()) *MockFloat64Observer_float64Observer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFloat64Observer_float64Observer_Call) Return() *MockFloat64Observer_float64Observer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFloat64Observer_float64Observer_Call) RunAndReturn(run func()) *MockFloat64Observer_float64Observer_Call {
	_c.Run(run)
	return _c
}

// NewMockFloat64Observer creates a new instance of MockFloat64Observer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64Observer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64Observer {
	mock := &MockFloat64Observer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
