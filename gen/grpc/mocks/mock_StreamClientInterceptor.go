// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockStreamClientInterceptor is an autogenerated mock type for the StreamClientInterceptor type
type MockStreamClientInterceptor struct {
	mock.Mock
}

type MockStreamClientInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamClientInterceptor) EXPECT() *MockStreamClientInterceptor_Expecter {
	return &MockStreamClientInterceptor_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, desc, cc, method, streamer, opts
func (_m *MockStreamClientInterceptor) Execute(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, desc, cc, method, streamer)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 grpc.ClientStream
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, grpc.Streamer, ...grpc.CallOption) (grpc.ClientStream, error)); ok {
		return rf(ctx, desc, cc, method, streamer, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, grpc.Streamer, ...grpc.CallOption) grpc.ClientStream); ok {
		r0 = rf(ctx, desc, cc, method, streamer, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grpc.ClientStream)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, grpc.Streamer, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, desc, cc, method, streamer, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamClientInterceptor_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockStreamClientInterceptor_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - desc *grpc.StreamDesc
//   - cc *grpc.ClientConn
//   - method string
//   - streamer grpc.Streamer
//   - opts ...grpc.CallOption
func (_e *MockStreamClientInterceptor_Expecter) Execute(ctx interface{}, desc interface{}, cc interface{}, method interface{}, streamer interface{}, opts ...interface{}) *MockStreamClientInterceptor_Execute_Call {
	return &MockStreamClientInterceptor_Execute_Call{Call: _e.mock.On("Execute",
		append([]interface{}{ctx, desc, cc, method, streamer}, opts...)...)}
}

func (_c *MockStreamClientInterceptor_Execute_Call) Run(run func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption)) *MockStreamClientInterceptor_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-5)
		for i, a := range args[5:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*grpc.StreamDesc), args[2].(*grpc.ClientConn), args[3].(string), args[4].(grpc.Streamer), variadicArgs...)
	})
	return _c
}

func (_c *MockStreamClientInterceptor_Execute_Call) Return(_a0 grpc.ClientStream, _a1 error) *MockStreamClientInterceptor_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamClientInterceptor_Execute_Call) RunAndReturn(run func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, grpc.Streamer, ...grpc.CallOption) (grpc.ClientStream, error)) *MockStreamClientInterceptor_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStreamClientInterceptor creates a new instance of MockStreamClientInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamClientInterceptor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamClientInterceptor {
	mock := &MockStreamClientInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
