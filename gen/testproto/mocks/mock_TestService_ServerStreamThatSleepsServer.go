// Code generated by mockery v2.52.3. DO NOT EDIT.

package mocks

import (
	context "context"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"

	testproto "github.com/neat-no/goat/gen/testproto"
)

// MockTestService_ServerStreamThatSleepsServer is an autogenerated mock type for the TestService_ServerStreamThatSleepsServer type
type MockTestService_ServerStreamThatSleepsServer[Res interface{}] struct {
	mock.Mock
}

type MockTestService_ServerStreamThatSleepsServer_Expecter[Res interface{}] struct {
	mock *mock.Mock
}

func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) EXPECT() *MockTestService_ServerStreamThatSleepsServer_Expecter[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_Expecter[Res]{mock: &_m.Mock}
}

// Context provides a mock function with no fields
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) Context() context.Context {
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

// MockTestService_ServerStreamThatSleepsServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockTestService_ServerStreamThatSleepsServer_Context_Call[Res interface{}] struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) Context() *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_Context_Call[Res]{Call: _e.mock.On("Context")}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res]) Run(run func()) *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res]) Return(_a0 context.Context) *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res]) RunAndReturn(run func() context.Context) *MockTestService_ServerStreamThatSleepsServer_Context_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) RecvMsg(m interface{}) error {
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

// MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res interface{}] struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) RecvMsg(m interface{}) *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res]{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res]) Run(run func(m interface{})) *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res]) Return(_a0 error) *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res]) RunAndReturn(run func(interface{}) error) *MockTestService_ServerStreamThatSleepsServer_RecvMsg_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) Send(_a0 *testproto.Msg) error {
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

// MockTestService_ServerStreamThatSleepsServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockTestService_ServerStreamThatSleepsServer_Send_Call[Res interface{}] struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *testproto.Msg
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) Send(_a0 interface{}) *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_Send_Call[Res]{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res]) Run(run func(_a0 *testproto.Msg)) *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*testproto.Msg))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res]) Return(_a0 error) *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res]) RunAndReturn(run func(*testproto.Msg) error) *MockTestService_ServerStreamThatSleepsServer_Send_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ServerStreamThatSleepsServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res interface{}] struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) SendHeader(_a0 interface{}) *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res]{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res]) Run(run func(_a0 metadata.MD)) *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res]) Return(_a0 error) *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res]) RunAndReturn(run func(metadata.MD) error) *MockTestService_ServerStreamThatSleepsServer_SendHeader_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) SendMsg(m interface{}) error {
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

// MockTestService_ServerStreamThatSleepsServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res interface{}] struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) SendMsg(m interface{}) *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res]{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res]) Run(run func(m interface{})) *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res]) Return(_a0 error) *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res]) RunAndReturn(run func(interface{}) error) *MockTestService_ServerStreamThatSleepsServer_SendMsg_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ServerStreamThatSleepsServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res interface{}] struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) SetHeader(_a0 interface{}) *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res]{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res]) Run(run func(_a0 metadata.MD)) *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res]) Return(_a0 error) *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res]) RunAndReturn(run func(metadata.MD) error) *MockTestService_ServerStreamThatSleepsServer_SetHeader_Call[Res] {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockTestService_ServerStreamThatSleepsServer[Res]) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res interface{}] struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ServerStreamThatSleepsServer_Expecter[Res]) SetTrailer(_a0 interface{}) *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res] {
	return &MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res]{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res]) Run(run func(_a0 metadata.MD)) *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res]) Return() *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res] {
	_c.Call.Return()
	return _c
}

func (_c *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res]) RunAndReturn(run func(metadata.MD)) *MockTestService_ServerStreamThatSleepsServer_SetTrailer_Call[Res] {
	_c.Run(run)
	return _c
}

// NewMockTestService_ServerStreamThatSleepsServer creates a new instance of MockTestService_ServerStreamThatSleepsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTestService_ServerStreamThatSleepsServer[Res interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTestService_ServerStreamThatSleepsServer[Res] {
	mock := &MockTestService_ServerStreamThatSleepsServer[Res]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
