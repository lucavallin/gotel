// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockFloat64Counter is an autogenerated mock type for the Float64Counter type
type MockFloat64Counter struct {
	mock.Mock
}

type MockFloat64Counter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64Counter) EXPECT() *MockFloat64Counter_Expecter {
	return &MockFloat64Counter_Expecter{mock: &_m.Mock}
}

// float64Counter provides a mock function with no fields
func (_m *MockFloat64Counter) float64Counter() {
	_m.Called()
}

// MockFloat64Counter_float64Counter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'float64Counter'
type MockFloat64Counter_float64Counter_Call struct {
	*mock.Call
}

// float64Counter is a helper method to define mock.On call
func (_e *MockFloat64Counter_Expecter) float64Counter() *MockFloat64Counter_float64Counter_Call {
	return &MockFloat64Counter_float64Counter_Call{Call: _e.mock.On("float64Counter")}
}

func (_c *MockFloat64Counter_float64Counter_Call) Run(run func()) *MockFloat64Counter_float64Counter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFloat64Counter_float64Counter_Call) Return() *MockFloat64Counter_float64Counter_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFloat64Counter_float64Counter_Call) RunAndReturn(run func()) *MockFloat64Counter_float64Counter_Call {
	_c.Run(run)
	return _c
}

// NewMockFloat64Counter creates a new instance of MockFloat64Counter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64Counter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64Counter {
	mock := &MockFloat64Counter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
