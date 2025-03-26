// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"

	testproto "github.com/neat-no/goat/gen/testproto"
)

// MockTestService_BidiStreamClient is an autogenerated mock type for the TestService_BidiStreamClient type
type MockTestService_BidiStreamClient[Req interface{}, Res interface{}] struct {
	mock.Mock
}

type MockTestService_BidiStreamClient_Expecter[Req interface{}, Res interface{}] struct {
	mock *mock.Mock
}

func (_m *MockTestService_BidiStreamClient[Req, Res]) EXPECT() *MockTestService_BidiStreamClient_Expecter[Req, Res] {
	return &MockTestService_BidiStreamClient_Expecter[Req, Res]{mock: &_m.Mock}
}

// CloseSend provides a mock function with no fields
func (_m *MockTestService_BidiStreamClient[Req, Res]) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_BidiStreamClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockTestService_BidiStreamClient_CloseSend_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) CloseSend() *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_CloseSend_Call[Req, Res]{Call: _e.mock.On("CloseSend")}
}

func (_c *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res]) Run(run func()) *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res]) Return(_a0 error) *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res]) RunAndReturn(run func() error) *MockTestService_BidiStreamClient_CloseSend_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with no fields
func (_m *MockTestService_BidiStreamClient[Req, Res]) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockTestService_BidiStreamClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockTestService_BidiStreamClient_Context_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) Context() *MockTestService_BidiStreamClient_Context_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_Context_Call[Req, Res]{Call: _e.mock.On("Context")}
}

func (_c *MockTestService_BidiStreamClient_Context_Call[Req, Res]) Run(run func()) *MockTestService_BidiStreamClient_Context_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_Context_Call[Req, Res]) Return(_a0 context.Context) *MockTestService_BidiStreamClient_Context_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_Context_Call[Req, Res]) RunAndReturn(run func() context.Context) *MockTestService_BidiStreamClient_Context_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with no fields
func (_m *MockTestService_BidiStreamClient[Req, Res]) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestService_BidiStreamClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockTestService_BidiStreamClient_Header_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) Header() *MockTestService_BidiStreamClient_Header_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_Header_Call[Req, Res]{Call: _e.mock.On("Header")}
}

func (_c *MockTestService_BidiStreamClient_Header_Call[Req, Res]) Run(run func()) *MockTestService_BidiStreamClient_Header_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_Header_Call[Req, Res]) Return(_a0 metadata.MD, _a1 error) *MockTestService_BidiStreamClient_Header_Call[Req, Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestService_BidiStreamClient_Header_Call[Req, Res]) RunAndReturn(run func() (metadata.MD, error)) *MockTestService_BidiStreamClient_Header_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with no fields
func (_m *MockTestService_BidiStreamClient[Req, Res]) Recv() (*testproto.Msg, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *testproto.Msg
	var r1 error
	if rf, ok := ret.Get(0).(func() (*testproto.Msg, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *testproto.Msg); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*testproto.Msg)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTestService_BidiStreamClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockTestService_BidiStreamClient_Recv_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) Recv() *MockTestService_BidiStreamClient_Recv_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_Recv_Call[Req, Res]{Call: _e.mock.On("Recv")}
}

func (_c *MockTestService_BidiStreamClient_Recv_Call[Req, Res]) Run(run func()) *MockTestService_BidiStreamClient_Recv_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_Recv_Call[Req, Res]) Return(_a0 *testproto.Msg, _a1 error) *MockTestService_BidiStreamClient_Recv_Call[Req, Res] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestService_BidiStreamClient_Recv_Call[Req, Res]) RunAndReturn(run func() (*testproto.Msg, error)) *MockTestService_BidiStreamClient_Recv_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockTestService_BidiStreamClient[Req, Res]) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_BidiStreamClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockTestService_BidiStreamClient_RecvMsg_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) RecvMsg(m interface{}) *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res]{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res]) Run(run func(m interface{})) *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res]) Return(_a0 error) *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res]) RunAndReturn(run func(interface{}) error) *MockTestService_BidiStreamClient_RecvMsg_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockTestService_BidiStreamClient[Req, Res]) Send(_a0 *testproto.Msg) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*testproto.Msg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_BidiStreamClient_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockTestService_BidiStreamClient_Send_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *testproto.Msg
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) Send(_a0 interface{}) *MockTestService_BidiStreamClient_Send_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_Send_Call[Req, Res]{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockTestService_BidiStreamClient_Send_Call[Req, Res]) Run(run func(_a0 *testproto.Msg)) *MockTestService_BidiStreamClient_Send_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*testproto.Msg))
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_Send_Call[Req, Res]) Return(_a0 error) *MockTestService_BidiStreamClient_Send_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_Send_Call[Req, Res]) RunAndReturn(run func(*testproto.Msg) error) *MockTestService_BidiStreamClient_Send_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockTestService_BidiStreamClient[Req, Res]) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_BidiStreamClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockTestService_BidiStreamClient_SendMsg_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) SendMsg(m interface{}) *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_SendMsg_Call[Req, Res]{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res]) Run(run func(m interface{})) *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res]) Return(_a0 error) *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res]) RunAndReturn(run func(interface{}) error) *MockTestService_BidiStreamClient_SendMsg_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with no fields
func (_m *MockTestService_BidiStreamClient[Req, Res]) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// MockTestService_BidiStreamClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockTestService_BidiStreamClient_Trailer_Call[Req interface{}, Res interface{}] struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockTestService_BidiStreamClient_Expecter[Req, Res]) Trailer() *MockTestService_BidiStreamClient_Trailer_Call[Req, Res] {
	return &MockTestService_BidiStreamClient_Trailer_Call[Req, Res]{Call: _e.mock.On("Trailer")}
}

func (_c *MockTestService_BidiStreamClient_Trailer_Call[Req, Res]) Run(run func()) *MockTestService_BidiStreamClient_Trailer_Call[Req, Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_BidiStreamClient_Trailer_Call[Req, Res]) Return(_a0 metadata.MD) *MockTestService_BidiStreamClient_Trailer_Call[Req, Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_BidiStreamClient_Trailer_Call[Req, Res]) RunAndReturn(run func() metadata.MD) *MockTestService_BidiStreamClient_Trailer_Call[Req, Res] {
	_c.Call.Return(run)
	return _c
}

// NewMockTestService_BidiStreamClient creates a new instance of MockTestService_BidiStreamClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTestService_BidiStreamClient[Req interface{}, Res interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTestService_BidiStreamClient[Req, Res] {
	mock := &MockTestService_BidiStreamClient[Req, Res]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
