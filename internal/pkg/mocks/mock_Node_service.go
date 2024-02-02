// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	node "github.com/awlsring/texit/internal/pkg/domain/node"
	mock "github.com/stretchr/testify/mock"
)

// MockNode_service is an autogenerated mock type for the Node type
type MockNode_service struct {
	mock.Mock
}

type MockNode_service_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNode_service) EXPECT() *MockNode_service_Expecter {
	return &MockNode_service_Expecter{mock: &_m.Mock}
}

// Describe provides a mock function with given fields: ctx, id
func (_m *MockNode_service) Describe(ctx context.Context, id node.Identifier) (*node.Node, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Describe")
	}

	var r0 *node.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) (*node.Node, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) *node.Node); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, node.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNode_service_Describe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Describe'
type MockNode_service_Describe_Call struct {
	*mock.Call
}

// Describe is a helper method to define mock.On call
//   - ctx context.Context
//   - id node.Identifier
func (_e *MockNode_service_Expecter) Describe(ctx interface{}, id interface{}) *MockNode_service_Describe_Call {
	return &MockNode_service_Describe_Call{Call: _e.mock.On("Describe", ctx, id)}
}

func (_c *MockNode_service_Describe_Call) Run(run func(ctx context.Context, id node.Identifier)) *MockNode_service_Describe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(node.Identifier))
	})
	return _c
}

func (_c *MockNode_service_Describe_Call) Return(_a0 *node.Node, _a1 error) *MockNode_service_Describe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNode_service_Describe_Call) RunAndReturn(run func(context.Context, node.Identifier) (*node.Node, error)) *MockNode_service_Describe_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockNode_service) List(ctx context.Context) ([]*node.Node, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*node.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*node.Node, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*node.Node); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*node.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNode_service_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockNode_service_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNode_service_Expecter) List(ctx interface{}) *MockNode_service_List_Call {
	return &MockNode_service_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockNode_service_List_Call) Run(run func(ctx context.Context)) *MockNode_service_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNode_service_List_Call) Return(_a0 []*node.Node, _a1 error) *MockNode_service_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNode_service_List_Call) RunAndReturn(run func(context.Context) ([]*node.Node, error)) *MockNode_service_List_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx, id
func (_m *MockNode_service) Start(ctx context.Context, id node.Identifier) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNode_service_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockNode_service_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
//   - id node.Identifier
func (_e *MockNode_service_Expecter) Start(ctx interface{}, id interface{}) *MockNode_service_Start_Call {
	return &MockNode_service_Start_Call{Call: _e.mock.On("Start", ctx, id)}
}

func (_c *MockNode_service_Start_Call) Run(run func(ctx context.Context, id node.Identifier)) *MockNode_service_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(node.Identifier))
	})
	return _c
}

func (_c *MockNode_service_Start_Call) Return(_a0 error) *MockNode_service_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNode_service_Start_Call) RunAndReturn(run func(context.Context, node.Identifier) error) *MockNode_service_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Status provides a mock function with given fields: ctx, id
func (_m *MockNode_service) Status(ctx context.Context, id node.Identifier) (node.Status, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 node.Status
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) (node.Status, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) node.Status); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(node.Status)
	}

	if rf, ok := ret.Get(1).(func(context.Context, node.Identifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNode_service_Status_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Status'
type MockNode_service_Status_Call struct {
	*mock.Call
}

// Status is a helper method to define mock.On call
//   - ctx context.Context
//   - id node.Identifier
func (_e *MockNode_service_Expecter) Status(ctx interface{}, id interface{}) *MockNode_service_Status_Call {
	return &MockNode_service_Status_Call{Call: _e.mock.On("Status", ctx, id)}
}

func (_c *MockNode_service_Status_Call) Run(run func(ctx context.Context, id node.Identifier)) *MockNode_service_Status_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(node.Identifier))
	})
	return _c
}

func (_c *MockNode_service_Status_Call) Return(_a0 node.Status, _a1 error) *MockNode_service_Status_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNode_service_Status_Call) RunAndReturn(run func(context.Context, node.Identifier) (node.Status, error)) *MockNode_service_Status_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: ctx, id
func (_m *MockNode_service) Stop(ctx context.Context, id node.Identifier) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, node.Identifier) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNode_service_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockNode_service_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - ctx context.Context
//   - id node.Identifier
func (_e *MockNode_service_Expecter) Stop(ctx interface{}, id interface{}) *MockNode_service_Stop_Call {
	return &MockNode_service_Stop_Call{Call: _e.mock.On("Stop", ctx, id)}
}

func (_c *MockNode_service_Stop_Call) Run(run func(ctx context.Context, id node.Identifier)) *MockNode_service_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(node.Identifier))
	})
	return _c
}

func (_c *MockNode_service_Stop_Call) Return(_a0 error) *MockNode_service_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNode_service_Stop_Call) RunAndReturn(run func(context.Context, node.Identifier) error) *MockNode_service_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNode_service creates a new instance of MockNode_service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNode_service(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNode_service {
	mock := &MockNode_service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}