// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockUnaryServerInterceptor is an autogenerated mock type for the UnaryServerInterceptor type
type MockUnaryServerInterceptor struct {
	mock.Mock
}

type MockUnaryServerInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnaryServerInterceptor) EXPECT() *MockUnaryServerInterceptor_Expecter {
	return &MockUnaryServerInterceptor_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, req, info, handler
func (_m *MockUnaryServerInterceptor) Execute(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ret := _m.Called(ctx, req, info, handler)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)); ok {
		return rf(ctx, req, info, handler)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) interface{}); ok {
		r0 = rf(ctx, req, info, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) error); ok {
		r1 = rf(ctx, req, info, handler)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUnaryServerInterceptor_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockUnaryServerInterceptor_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - req interface{}
//   - info *grpc.UnaryServerInfo
//   - handler grpc.UnaryHandler
func (_e *MockUnaryServerInterceptor_Expecter) Execute(ctx interface{}, req interface{}, info interface{}, handler interface{}) *MockUnaryServerInterceptor_Execute_Call {
	return &MockUnaryServerInterceptor_Execute_Call{Call: _e.mock.On("Execute", ctx, req, info, handler)}
}

func (_c *MockUnaryServerInterceptor_Execute_Call) Run(run func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)) *MockUnaryServerInterceptor_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}), args[2].(*grpc.UnaryServerInfo), args[3].(grpc.UnaryHandler))
	})
	return _c
}

func (_c *MockUnaryServerInterceptor_Execute_Call) Return(resp interface{}, err error) *MockUnaryServerInterceptor_Execute_Call {
	_c.Call.Return(resp, err)
	return _c
}

func (_c *MockUnaryServerInterceptor_Execute_Call) RunAndReturn(run func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error)) *MockUnaryServerInterceptor_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUnaryServerInterceptor creates a new instance of MockUnaryServerInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnaryServerInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnaryServerInterceptor {
	mock := &MockUnaryServerInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
