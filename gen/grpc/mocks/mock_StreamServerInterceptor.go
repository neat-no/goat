// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"
)

// MockStreamServerInterceptor is an autogenerated mock type for the StreamServerInterceptor type
type MockStreamServerInterceptor struct {
	mock.Mock
}

type MockStreamServerInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamServerInterceptor) EXPECT() *MockStreamServerInterceptor_Expecter {
	return &MockStreamServerInterceptor_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: srv, ss, info, handler
func (_m *MockStreamServerInterceptor) Execute(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ret := _m.Called(srv, ss, info, handler)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error); ok {
		r0 = rf(srv, ss, info, handler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamServerInterceptor_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockStreamServerInterceptor_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - srv interface{}
//   - ss grpc.ServerStream
//   - info *grpc.StreamServerInfo
//   - handler grpc.StreamHandler
func (_e *MockStreamServerInterceptor_Expecter) Execute(srv interface{}, ss interface{}, info interface{}, handler interface{}) *MockStreamServerInterceptor_Execute_Call {
	return &MockStreamServerInterceptor_Execute_Call{Call: _e.mock.On("Execute", srv, ss, info, handler)}
}

func (_c *MockStreamServerInterceptor_Execute_Call) Run(run func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler)) *MockStreamServerInterceptor_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}), args[1].(grpc.ServerStream), args[2].(*grpc.StreamServerInfo), args[3].(grpc.StreamHandler))
	})
	return _c
}

func (_c *MockStreamServerInterceptor_Execute_Call) Return(_a0 error) *MockStreamServerInterceptor_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamServerInterceptor_Execute_Call) RunAndReturn(run func(interface{}, grpc.ServerStream, *grpc.StreamServerInfo, grpc.StreamHandler) error) *MockStreamServerInterceptor_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStreamServerInterceptor creates a new instance of MockStreamServerInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamServerInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamServerInterceptor {
	mock := &MockStreamServerInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
