// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	linodego "github.com/linode/linodego"

	mock "github.com/stretchr/testify/mock"
)

// MockLinodeClient_interfaces is an autogenerated mock type for the LinodeClient type
type MockLinodeClient_interfaces struct {
	mock.Mock
}

type MockLinodeClient_interfaces_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLinodeClient_interfaces) EXPECT() *MockLinodeClient_interfaces_Expecter {
	return &MockLinodeClient_interfaces_Expecter{mock: &_m.Mock}
}

// BootInstance provides a mock function with given fields: ctx, linodeID, configID
func (_m *MockLinodeClient_interfaces) BootInstance(ctx context.Context, linodeID int, configID int) error {
	ret := _m.Called(ctx, linodeID, configID)

	if len(ret) == 0 {
		panic("no return value specified for BootInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, linodeID, configID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLinodeClient_interfaces_BootInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BootInstance'
type MockLinodeClient_interfaces_BootInstance_Call struct {
	*mock.Call
}

// BootInstance is a helper method to define mock.On call
//   - ctx context.Context
//   - linodeID int
//   - configID int
func (_e *MockLinodeClient_interfaces_Expecter) BootInstance(ctx interface{}, linodeID interface{}, configID interface{}) *MockLinodeClient_interfaces_BootInstance_Call {
	return &MockLinodeClient_interfaces_BootInstance_Call{Call: _e.mock.On("BootInstance", ctx, linodeID, configID)}
}

func (_c *MockLinodeClient_interfaces_BootInstance_Call) Run(run func(ctx context.Context, linodeID int, configID int)) *MockLinodeClient_interfaces_BootInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_BootInstance_Call) Return(_a0 error) *MockLinodeClient_interfaces_BootInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLinodeClient_interfaces_BootInstance_Call) RunAndReturn(run func(context.Context, int, int) error) *MockLinodeClient_interfaces_BootInstance_Call {
	_c.Call.Return(run)
	return _c
}

// CreateInstance provides a mock function with given fields: ctx, opts
func (_m *MockLinodeClient_interfaces) CreateInstance(ctx context.Context, opts linodego.InstanceCreateOptions) (*linodego.Instance, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for CreateInstance")
	}

	var r0 *linodego.Instance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, linodego.InstanceCreateOptions) (*linodego.Instance, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, linodego.InstanceCreateOptions) *linodego.Instance); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*linodego.Instance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, linodego.InstanceCreateOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLinodeClient_interfaces_CreateInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInstance'
type MockLinodeClient_interfaces_CreateInstance_Call struct {
	*mock.Call
}

// CreateInstance is a helper method to define mock.On call
//   - ctx context.Context
//   - opts linodego.InstanceCreateOptions
func (_e *MockLinodeClient_interfaces_Expecter) CreateInstance(ctx interface{}, opts interface{}) *MockLinodeClient_interfaces_CreateInstance_Call {
	return &MockLinodeClient_interfaces_CreateInstance_Call{Call: _e.mock.On("CreateInstance", ctx, opts)}
}

func (_c *MockLinodeClient_interfaces_CreateInstance_Call) Run(run func(ctx context.Context, opts linodego.InstanceCreateOptions)) *MockLinodeClient_interfaces_CreateInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(linodego.InstanceCreateOptions))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_CreateInstance_Call) Return(_a0 *linodego.Instance, _a1 error) *MockLinodeClient_interfaces_CreateInstance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLinodeClient_interfaces_CreateInstance_Call) RunAndReturn(run func(context.Context, linodego.InstanceCreateOptions) (*linodego.Instance, error)) *MockLinodeClient_interfaces_CreateInstance_Call {
	_c.Call.Return(run)
	return _c
}

// CreateStackscript provides a mock function with given fields: ctx, opts
func (_m *MockLinodeClient_interfaces) CreateStackscript(ctx context.Context, opts linodego.StackscriptCreateOptions) (*linodego.Stackscript, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for CreateStackscript")
	}

	var r0 *linodego.Stackscript
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, linodego.StackscriptCreateOptions) (*linodego.Stackscript, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, linodego.StackscriptCreateOptions) *linodego.Stackscript); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*linodego.Stackscript)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, linodego.StackscriptCreateOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLinodeClient_interfaces_CreateStackscript_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateStackscript'
type MockLinodeClient_interfaces_CreateStackscript_Call struct {
	*mock.Call
}

// CreateStackscript is a helper method to define mock.On call
//   - ctx context.Context
//   - opts linodego.StackscriptCreateOptions
func (_e *MockLinodeClient_interfaces_Expecter) CreateStackscript(ctx interface{}, opts interface{}) *MockLinodeClient_interfaces_CreateStackscript_Call {
	return &MockLinodeClient_interfaces_CreateStackscript_Call{Call: _e.mock.On("CreateStackscript", ctx, opts)}
}

func (_c *MockLinodeClient_interfaces_CreateStackscript_Call) Run(run func(ctx context.Context, opts linodego.StackscriptCreateOptions)) *MockLinodeClient_interfaces_CreateStackscript_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(linodego.StackscriptCreateOptions))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_CreateStackscript_Call) Return(_a0 *linodego.Stackscript, _a1 error) *MockLinodeClient_interfaces_CreateStackscript_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLinodeClient_interfaces_CreateStackscript_Call) RunAndReturn(run func(context.Context, linodego.StackscriptCreateOptions) (*linodego.Stackscript, error)) *MockLinodeClient_interfaces_CreateStackscript_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteInstance provides a mock function with given fields: ctx, linodeID
func (_m *MockLinodeClient_interfaces) DeleteInstance(ctx context.Context, linodeID int) error {
	ret := _m.Called(ctx, linodeID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, linodeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLinodeClient_interfaces_DeleteInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteInstance'
type MockLinodeClient_interfaces_DeleteInstance_Call struct {
	*mock.Call
}

// DeleteInstance is a helper method to define mock.On call
//   - ctx context.Context
//   - linodeID int
func (_e *MockLinodeClient_interfaces_Expecter) DeleteInstance(ctx interface{}, linodeID interface{}) *MockLinodeClient_interfaces_DeleteInstance_Call {
	return &MockLinodeClient_interfaces_DeleteInstance_Call{Call: _e.mock.On("DeleteInstance", ctx, linodeID)}
}

func (_c *MockLinodeClient_interfaces_DeleteInstance_Call) Run(run func(ctx context.Context, linodeID int)) *MockLinodeClient_interfaces_DeleteInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_DeleteInstance_Call) Return(_a0 error) *MockLinodeClient_interfaces_DeleteInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLinodeClient_interfaces_DeleteInstance_Call) RunAndReturn(run func(context.Context, int) error) *MockLinodeClient_interfaces_DeleteInstance_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteStackscript provides a mock function with given fields: ctx, scriptID
func (_m *MockLinodeClient_interfaces) DeleteStackscript(ctx context.Context, scriptID int) error {
	ret := _m.Called(ctx, scriptID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteStackscript")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, scriptID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLinodeClient_interfaces_DeleteStackscript_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteStackscript'
type MockLinodeClient_interfaces_DeleteStackscript_Call struct {
	*mock.Call
}

// DeleteStackscript is a helper method to define mock.On call
//   - ctx context.Context
//   - scriptID int
func (_e *MockLinodeClient_interfaces_Expecter) DeleteStackscript(ctx interface{}, scriptID interface{}) *MockLinodeClient_interfaces_DeleteStackscript_Call {
	return &MockLinodeClient_interfaces_DeleteStackscript_Call{Call: _e.mock.On("DeleteStackscript", ctx, scriptID)}
}

func (_c *MockLinodeClient_interfaces_DeleteStackscript_Call) Run(run func(ctx context.Context, scriptID int)) *MockLinodeClient_interfaces_DeleteStackscript_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_DeleteStackscript_Call) Return(_a0 error) *MockLinodeClient_interfaces_DeleteStackscript_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLinodeClient_interfaces_DeleteStackscript_Call) RunAndReturn(run func(context.Context, int) error) *MockLinodeClient_interfaces_DeleteStackscript_Call {
	_c.Call.Return(run)
	return _c
}

// GetInstance provides a mock function with given fields: ctx, linodeID
func (_m *MockLinodeClient_interfaces) GetInstance(ctx context.Context, linodeID int) (*linodego.Instance, error) {
	ret := _m.Called(ctx, linodeID)

	if len(ret) == 0 {
		panic("no return value specified for GetInstance")
	}

	var r0 *linodego.Instance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*linodego.Instance, error)); ok {
		return rf(ctx, linodeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *linodego.Instance); ok {
		r0 = rf(ctx, linodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*linodego.Instance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, linodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLinodeClient_interfaces_GetInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInstance'
type MockLinodeClient_interfaces_GetInstance_Call struct {
	*mock.Call
}

// GetInstance is a helper method to define mock.On call
//   - ctx context.Context
//   - linodeID int
func (_e *MockLinodeClient_interfaces_Expecter) GetInstance(ctx interface{}, linodeID interface{}) *MockLinodeClient_interfaces_GetInstance_Call {
	return &MockLinodeClient_interfaces_GetInstance_Call{Call: _e.mock.On("GetInstance", ctx, linodeID)}
}

func (_c *MockLinodeClient_interfaces_GetInstance_Call) Run(run func(ctx context.Context, linodeID int)) *MockLinodeClient_interfaces_GetInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_GetInstance_Call) Return(_a0 *linodego.Instance, _a1 error) *MockLinodeClient_interfaces_GetInstance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLinodeClient_interfaces_GetInstance_Call) RunAndReturn(run func(context.Context, int) (*linodego.Instance, error)) *MockLinodeClient_interfaces_GetInstance_Call {
	_c.Call.Return(run)
	return _c
}

// ListStackscripts provides a mock function with given fields: ctx, opts
func (_m *MockLinodeClient_interfaces) ListStackscripts(ctx context.Context, opts *linodego.ListOptions) ([]linodego.Stackscript, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for ListStackscripts")
	}

	var r0 []linodego.Stackscript
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *linodego.ListOptions) ([]linodego.Stackscript, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *linodego.ListOptions) []linodego.Stackscript); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]linodego.Stackscript)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *linodego.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLinodeClient_interfaces_ListStackscripts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListStackscripts'
type MockLinodeClient_interfaces_ListStackscripts_Call struct {
	*mock.Call
}

// ListStackscripts is a helper method to define mock.On call
//   - ctx context.Context
//   - opts *linodego.ListOptions
func (_e *MockLinodeClient_interfaces_Expecter) ListStackscripts(ctx interface{}, opts interface{}) *MockLinodeClient_interfaces_ListStackscripts_Call {
	return &MockLinodeClient_interfaces_ListStackscripts_Call{Call: _e.mock.On("ListStackscripts", ctx, opts)}
}

func (_c *MockLinodeClient_interfaces_ListStackscripts_Call) Run(run func(ctx context.Context, opts *linodego.ListOptions)) *MockLinodeClient_interfaces_ListStackscripts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*linodego.ListOptions))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_ListStackscripts_Call) Return(_a0 []linodego.Stackscript, _a1 error) *MockLinodeClient_interfaces_ListStackscripts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLinodeClient_interfaces_ListStackscripts_Call) RunAndReturn(run func(context.Context, *linodego.ListOptions) ([]linodego.Stackscript, error)) *MockLinodeClient_interfaces_ListStackscripts_Call {
	_c.Call.Return(run)
	return _c
}

// ShutdownInstance provides a mock function with given fields: ctx, id
func (_m *MockLinodeClient_interfaces) ShutdownInstance(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ShutdownInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLinodeClient_interfaces_ShutdownInstance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShutdownInstance'
type MockLinodeClient_interfaces_ShutdownInstance_Call struct {
	*mock.Call
}

// ShutdownInstance is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
func (_e *MockLinodeClient_interfaces_Expecter) ShutdownInstance(ctx interface{}, id interface{}) *MockLinodeClient_interfaces_ShutdownInstance_Call {
	return &MockLinodeClient_interfaces_ShutdownInstance_Call{Call: _e.mock.On("ShutdownInstance", ctx, id)}
}

func (_c *MockLinodeClient_interfaces_ShutdownInstance_Call) Run(run func(ctx context.Context, id int)) *MockLinodeClient_interfaces_ShutdownInstance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockLinodeClient_interfaces_ShutdownInstance_Call) Return(_a0 error) *MockLinodeClient_interfaces_ShutdownInstance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLinodeClient_interfaces_ShutdownInstance_Call) RunAndReturn(run func(context.Context, int) error) *MockLinodeClient_interfaces_ShutdownInstance_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLinodeClient_interfaces creates a new instance of MockLinodeClient_interfaces. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLinodeClient_interfaces(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLinodeClient_interfaces {
	mock := &MockLinodeClient_interfaces{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}