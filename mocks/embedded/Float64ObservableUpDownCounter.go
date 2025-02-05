// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockFloat64ObservableUpDownCounter is an autogenerated mock type for the Float64ObservableUpDownCounter type
type MockFloat64ObservableUpDownCounter struct {
	mock.Mock
}

type MockFloat64ObservableUpDownCounter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64ObservableUpDownCounter) EXPECT() *MockFloat64ObservableUpDownCounter_Expecter {
	return &MockFloat64ObservableUpDownCounter_Expecter{mock: &_m.Mock}
}

// float64ObservableUpDownCounter provides a mock function with no fields
func (_m *MockFloat64ObservableUpDownCounter) float64ObservableUpDownCounter() {
	_m.Called()
}

// MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'float64ObservableUpDownCounter'
type MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call struct {
	*mock.Call
}

// float64ObservableUpDownCounter is a helper method to define mock.On call
func (_e *MockFloat64ObservableUpDownCounter_Expecter) float64ObservableUpDownCounter() *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call {
	return &MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call{Call: _e.mock.On("float64ObservableUpDownCounter")}
}

func (_c *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call) Run(run func()) *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call) Return() *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call) RunAndReturn(run func()) *MockFloat64ObservableUpDownCounter_float64ObservableUpDownCounter_Call {
	_c.Run(run)
	return _c
}

// NewMockFloat64ObservableUpDownCounter creates a new instance of MockFloat64ObservableUpDownCounter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64ObservableUpDownCounter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64ObservableUpDownCounter {
	mock := &MockFloat64ObservableUpDownCounter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
