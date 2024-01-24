// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	testproto "github.com/avos-io/goat/gen/testproto"
)

// MockTestService_ClientStreamServer is an autogenerated mock type for the TestService_ClientStreamServer type
type MockTestService_ClientStreamServer struct {
	mock.Mock
}

type MockTestService_ClientStreamServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTestService_ClientStreamServer) EXPECT() *MockTestService_ClientStreamServer_Expecter {
	return &MockTestService_ClientStreamServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockTestService_ClientStreamServer) Context() context.Context {
	ret := _m.Called()

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

// MockTestService_ClientStreamServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockTestService_ClientStreamServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockTestService_ClientStreamServer_Expecter) Context() *MockTestService_ClientStreamServer_Context_Call {
	return &MockTestService_ClientStreamServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockTestService_ClientStreamServer_Context_Call) Run(run func()) *MockTestService_ClientStreamServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_Context_Call) Return(_a0 context.Context) *MockTestService_ClientStreamServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_Context_Call) RunAndReturn(run func() context.Context) *MockTestService_ClientStreamServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *MockTestService_ClientStreamServer) Recv() (*testproto.Msg, error) {
	ret := _m.Called()

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

// MockTestService_ClientStreamServer_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockTestService_ClientStreamServer_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockTestService_ClientStreamServer_Expecter) Recv() *MockTestService_ClientStreamServer_Recv_Call {
	return &MockTestService_ClientStreamServer_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockTestService_ClientStreamServer_Recv_Call) Run(run func()) *MockTestService_ClientStreamServer_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_Recv_Call) Return(_a0 *testproto.Msg, _a1 error) *MockTestService_ClientStreamServer_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTestService_ClientStreamServer_Recv_Call) RunAndReturn(run func() (*testproto.Msg, error)) *MockTestService_ClientStreamServer_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockTestService_ClientStreamServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ClientStreamServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockTestService_ClientStreamServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_ClientStreamServer_Expecter) RecvMsg(m interface{}) *MockTestService_ClientStreamServer_RecvMsg_Call {
	return &MockTestService_ClientStreamServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockTestService_ClientStreamServer_RecvMsg_Call) Run(run func(m interface{})) *MockTestService_ClientStreamServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_RecvMsg_Call) Return(_a0 error) *MockTestService_ClientStreamServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockTestService_ClientStreamServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendAndClose provides a mock function with given fields: _a0
func (_m *MockTestService_ClientStreamServer) SendAndClose(_a0 *testproto.Msg) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*testproto.Msg) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ClientStreamServer_SendAndClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendAndClose'
type MockTestService_ClientStreamServer_SendAndClose_Call struct {
	*mock.Call
}

// SendAndClose is a helper method to define mock.On call
//   - _a0 *testproto.Msg
func (_e *MockTestService_ClientStreamServer_Expecter) SendAndClose(_a0 interface{}) *MockTestService_ClientStreamServer_SendAndClose_Call {
	return &MockTestService_ClientStreamServer_SendAndClose_Call{Call: _e.mock.On("SendAndClose", _a0)}
}

func (_c *MockTestService_ClientStreamServer_SendAndClose_Call) Run(run func(_a0 *testproto.Msg)) *MockTestService_ClientStreamServer_SendAndClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*testproto.Msg))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendAndClose_Call) Return(_a0 error) *MockTestService_ClientStreamServer_SendAndClose_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendAndClose_Call) RunAndReturn(run func(*testproto.Msg) error) *MockTestService_ClientStreamServer_SendAndClose_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockTestService_ClientStreamServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ClientStreamServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockTestService_ClientStreamServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ClientStreamServer_Expecter) SendHeader(_a0 interface{}) *MockTestService_ClientStreamServer_SendHeader_Call {
	return &MockTestService_ClientStreamServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockTestService_ClientStreamServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockTestService_ClientStreamServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendHeader_Call) Return(_a0 error) *MockTestService_ClientStreamServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockTestService_ClientStreamServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockTestService_ClientStreamServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ClientStreamServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockTestService_ClientStreamServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockTestService_ClientStreamServer_Expecter) SendMsg(m interface{}) *MockTestService_ClientStreamServer_SendMsg_Call {
	return &MockTestService_ClientStreamServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockTestService_ClientStreamServer_SendMsg_Call) Run(run func(m interface{})) *MockTestService_ClientStreamServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendMsg_Call) Return(_a0 error) *MockTestService_ClientStreamServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockTestService_ClientStreamServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockTestService_ClientStreamServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTestService_ClientStreamServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockTestService_ClientStreamServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ClientStreamServer_Expecter) SetHeader(_a0 interface{}) *MockTestService_ClientStreamServer_SetHeader_Call {
	return &MockTestService_ClientStreamServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockTestService_ClientStreamServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockTestService_ClientStreamServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_SetHeader_Call) Return(_a0 error) *MockTestService_ClientStreamServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTestService_ClientStreamServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockTestService_ClientStreamServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockTestService_ClientStreamServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockTestService_ClientStreamServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockTestService_ClientStreamServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockTestService_ClientStreamServer_Expecter) SetTrailer(_a0 interface{}) *MockTestService_ClientStreamServer_SetTrailer_Call {
	return &MockTestService_ClientStreamServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockTestService_ClientStreamServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockTestService_ClientStreamServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockTestService_ClientStreamServer_SetTrailer_Call) Return() *MockTestService_ClientStreamServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockTestService_ClientStreamServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *MockTestService_ClientStreamServer_SetTrailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTestService_ClientStreamServer creates a new instance of MockTestService_ClientStreamServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTestService_ClientStreamServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTestService_ClientStreamServer {
	mock := &MockTestService_ClientStreamServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
