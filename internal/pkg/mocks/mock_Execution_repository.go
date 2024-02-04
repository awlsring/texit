// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	workflow "github.com/awlsring/texit/internal/pkg/domain/workflow"
)

// MockExecution_repository is an autogenerated mock type for the Execution type
type MockExecution_repository struct {
	mock.Mock
}

type MockExecution_repository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExecution_repository) EXPECT() *MockExecution_repository_Expecter {
	return &MockExecution_repository_Expecter{mock: &_m.Mock}
}

// CloseExecution provides a mock function with given fields: ctx, id, result, msg
func (_m *MockExecution_repository) CloseExecution(ctx context.Context, id workflow.ExecutionIdentifier, result workflow.Status, msg []string) error {
	ret := _m.Called(ctx, id, result, msg)

	if len(ret) == 0 {
		panic("no return value specified for CloseExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workflow.ExecutionIdentifier, workflow.Status, []string) error); ok {
		r0 = rf(ctx, id, result, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExecution_repository_CloseExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseExecution'
type MockExecution_repository_CloseExecution_Call struct {
	*mock.Call
}

// CloseExecution is a helper method to define mock.On call
//   - ctx context.Context
//   - id workflow.ExecutionIdentifier
//   - result workflow.Status
//   - msg []string
func (_e *MockExecution_repository_Expecter) CloseExecution(ctx interface{}, id interface{}, result interface{}, msg interface{}) *MockExecution_repository_CloseExecution_Call {
	return &MockExecution_repository_CloseExecution_Call{Call: _e.mock.On("CloseExecution", ctx, id, result, msg)}
}

func (_c *MockExecution_repository_CloseExecution_Call) Run(run func(ctx context.Context, id workflow.ExecutionIdentifier, result workflow.Status, msg []string)) *MockExecution_repository_CloseExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workflow.ExecutionIdentifier), args[2].(workflow.Status), args[3].([]string))
	})
	return _c
}

func (_c *MockExecution_repository_CloseExecution_Call) Return(_a0 error) *MockExecution_repository_CloseExecution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecution_repository_CloseExecution_Call) RunAndReturn(run func(context.Context, workflow.ExecutionIdentifier, workflow.Status, []string) error) *MockExecution_repository_CloseExecution_Call {
	_c.Call.Return(run)
	return _c
}

// CreateExecution provides a mock function with given fields: ctx, ex
func (_m *MockExecution_repository) CreateExecution(ctx context.Context, ex *workflow.Execution) error {
	ret := _m.Called(ctx, ex)

	if len(ret) == 0 {
		panic("no return value specified for CreateExecution")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *workflow.Execution) error); ok {
		r0 = rf(ctx, ex)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExecution_repository_CreateExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateExecution'
type MockExecution_repository_CreateExecution_Call struct {
	*mock.Call
}

// CreateExecution is a helper method to define mock.On call
//   - ctx context.Context
//   - ex *workflow.Execution
func (_e *MockExecution_repository_Expecter) CreateExecution(ctx interface{}, ex interface{}) *MockExecution_repository_CreateExecution_Call {
	return &MockExecution_repository_CreateExecution_Call{Call: _e.mock.On("CreateExecution", ctx, ex)}
}

func (_c *MockExecution_repository_CreateExecution_Call) Run(run func(ctx context.Context, ex *workflow.Execution)) *MockExecution_repository_CreateExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*workflow.Execution))
	})
	return _c
}

func (_c *MockExecution_repository_CreateExecution_Call) Return(_a0 error) *MockExecution_repository_CreateExecution_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecution_repository_CreateExecution_Call) RunAndReturn(run func(context.Context, *workflow.Execution) error) *MockExecution_repository_CreateExecution_Call {
	_c.Call.Return(run)
	return _c
}

// GetExecution provides a mock function with given fields: ctx, id
func (_m *MockExecution_repository) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetExecution")
	}

	var r0 *workflow.Execution
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workflow.ExecutionIdentifier) *workflow.Execution); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workflow.Execution)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workflow.ExecutionIdentifier) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExecution_repository_GetExecution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExecution'
type MockExecution_repository_GetExecution_Call struct {
	*mock.Call
}

// GetExecution is a helper method to define mock.On call
//   - ctx context.Context
//   - id workflow.ExecutionIdentifier
func (_e *MockExecution_repository_Expecter) GetExecution(ctx interface{}, id interface{}) *MockExecution_repository_GetExecution_Call {
	return &MockExecution_repository_GetExecution_Call{Call: _e.mock.On("GetExecution", ctx, id)}
}

func (_c *MockExecution_repository_GetExecution_Call) Run(run func(ctx context.Context, id workflow.ExecutionIdentifier)) *MockExecution_repository_GetExecution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workflow.ExecutionIdentifier))
	})
	return _c
}

func (_c *MockExecution_repository_GetExecution_Call) Return(_a0 *workflow.Execution, _a1 error) *MockExecution_repository_GetExecution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExecution_repository_GetExecution_Call) RunAndReturn(run func(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)) *MockExecution_repository_GetExecution_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: ctx
func (_m *MockExecution_repository) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockExecution_repository_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockExecution_repository_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockExecution_repository_Expecter) Init(ctx interface{}) *MockExecution_repository_Init_Call {
	return &MockExecution_repository_Init_Call{Call: _e.mock.On("Init", ctx)}
}

func (_c *MockExecution_repository_Init_Call) Run(run func(ctx context.Context)) *MockExecution_repository_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockExecution_repository_Init_Call) Return(_a0 error) *MockExecution_repository_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockExecution_repository_Init_Call) RunAndReturn(run func(context.Context) error) *MockExecution_repository_Init_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExecution_repository creates a new instance of MockExecution_repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecution_repository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExecution_repository {
	mock := &MockExecution_repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}