// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockUnaryClientInterceptor is an autogenerated mock type for the UnaryClientInterceptor type
type MockUnaryClientInterceptor struct {
	mock.Mock
}

type MockUnaryClientInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnaryClientInterceptor) EXPECT() *MockUnaryClientInterceptor_Expecter {
	return &MockUnaryClientInterceptor_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, method, req, reply, cc, invoker, opts
func (_m *MockUnaryClientInterceptor) Execute(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, method, req, reply, cc, invoker)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, interface{}, *grpc.ClientConn, grpc.UnaryInvoker, ...grpc.CallOption) error); ok {
		r0 = rf(ctx, method, req, reply, cc, invoker, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUnaryClientInterceptor_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockUnaryClientInterceptor_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - method string
//   - req interface{}
//   - reply interface{}
//   - cc *grpc.ClientConn
//   - invoker grpc.UnaryInvoker
//   - opts ...grpc.CallOption
func (_e *MockUnaryClientInterceptor_Expecter) Execute(ctx interface{}, method interface{}, req interface{}, reply interface{}, cc interface{}, invoker interface{}, opts ...interface{}) *MockUnaryClientInterceptor_Execute_Call {
	return &MockUnaryClientInterceptor_Execute_Call{Call: _e.mock.On("Execute",
		append([]interface{}{ctx, method, req, reply, cc, invoker}, opts...)...)}
}

func (_c *MockUnaryClientInterceptor_Execute_Call) Run(run func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption)) *MockUnaryClientInterceptor_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-6)
		for i, a := range args[6:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(interface{}), args[4].(*grpc.ClientConn), args[5].(grpc.UnaryInvoker), variadicArgs...)
	})
	return _c
}

func (_c *MockUnaryClientInterceptor_Execute_Call) Return(_a0 error) *MockUnaryClientInterceptor_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUnaryClientInterceptor_Execute_Call) RunAndReturn(run func(context.Context, string, interface{}, interface{}, *grpc.ClientConn, grpc.UnaryInvoker, ...grpc.CallOption) error) *MockUnaryClientInterceptor_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUnaryClientInterceptor creates a new instance of MockUnaryClientInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnaryClientInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnaryClientInterceptor {
	mock := &MockUnaryClientInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
