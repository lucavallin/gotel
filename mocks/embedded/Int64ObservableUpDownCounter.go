// Code generated by mockery v2.50.0. DO NOT EDIT.

package embedded

import mock "github.com/stretchr/testify/mock"

// MockInt64ObservableUpDownCounter is an autogenerated mock type for the Int64ObservableUpDownCounter type
type MockInt64ObservableUpDownCounter struct {
	mock.Mock
}

type MockInt64ObservableUpDownCounter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInt64ObservableUpDownCounter) EXPECT() *MockInt64ObservableUpDownCounter_Expecter {
	return &MockInt64ObservableUpDownCounter_Expecter{mock: &_m.Mock}
}

// int64ObservableUpDownCounter provides a mock function with no fields
func (_m *MockInt64ObservableUpDownCounter) int64ObservableUpDownCounter() {
	_m.Called()
}

// MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'int64ObservableUpDownCounter'
type MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call struct {
	*mock.Call
}

// int64ObservableUpDownCounter is a helper method to define mock.On call
func (_e *MockInt64ObservableUpDownCounter_Expecter) int64ObservableUpDownCounter() *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call {
	return &MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call{Call: _e.mock.On("int64ObservableUpDownCounter")}
}

func (_c *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call) Run(run func()) *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call) Return() *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call) RunAndReturn(run func()) *MockInt64ObservableUpDownCounter_int64ObservableUpDownCounter_Call {
	_c.Run(run)
	return _c
}

// NewMockInt64ObservableUpDownCounter creates a new instance of MockInt64ObservableUpDownCounter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInt64ObservableUpDownCounter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInt64ObservableUpDownCounter {
	mock := &MockInt64ObservableUpDownCounter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
