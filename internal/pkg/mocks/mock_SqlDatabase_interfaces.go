// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// MockSqlDatabase_interfaces is an autogenerated mock type for the SqlDatabase type
type MockSqlDatabase_interfaces struct {
	mock.Mock
}

type MockSqlDatabase_interfaces_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSqlDatabase_interfaces) EXPECT() *MockSqlDatabase_interfaces_Expecter {
	return &MockSqlDatabase_interfaces_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockSqlDatabase_interfaces) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlDatabase_interfaces_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSqlDatabase_interfaces_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSqlDatabase_interfaces_Expecter) Close() *MockSqlDatabase_interfaces_Close_Call {
	return &MockSqlDatabase_interfaces_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSqlDatabase_interfaces_Close_Call) Run(run func()) *MockSqlDatabase_interfaces_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSqlDatabase_interfaces_Close_Call) Return(_a0 error) *MockSqlDatabase_interfaces_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlDatabase_interfaces_Close_Call) RunAndReturn(run func() error) *MockSqlDatabase_interfaces_Close_Call {
	_c.Call.Return(run)
	return _c
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *MockSqlDatabase_interfaces) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_interfaces_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type MockSqlDatabase_interfaces_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockSqlDatabase_interfaces_Expecter) ExecContext(ctx interface{}, query interface{}, args ...interface{}) *MockSqlDatabase_interfaces_ExecContext_Call {
	return &MockSqlDatabase_interfaces_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockSqlDatabase_interfaces_ExecContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockSqlDatabase_interfaces_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSqlDatabase_interfaces_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockSqlDatabase_interfaces_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_interfaces_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *MockSqlDatabase_interfaces_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockSqlDatabase_interfaces) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlDatabase_interfaces_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSqlDatabase_interfaces_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockSqlDatabase_interfaces_Expecter) GetContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockSqlDatabase_interfaces_GetContext_Call {
	return &MockSqlDatabase_interfaces_GetContext_Call{Call: _e.mock.On("GetContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockSqlDatabase_interfaces_GetContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockSqlDatabase_interfaces_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSqlDatabase_interfaces_GetContext_Call) Return(_a0 error) *MockSqlDatabase_interfaces_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlDatabase_interfaces_GetContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockSqlDatabase_interfaces_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// SelectContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockSqlDatabase_interfaces) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlDatabase_interfaces_SelectContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectContext'
type MockSqlDatabase_interfaces_SelectContext_Call struct {
	*mock.Call
}

// SelectContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockSqlDatabase_interfaces_Expecter) SelectContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockSqlDatabase_interfaces_SelectContext_Call {
	return &MockSqlDatabase_interfaces_SelectContext_Call{Call: _e.mock.On("SelectContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockSqlDatabase_interfaces_SelectContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockSqlDatabase_interfaces_SelectContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockSqlDatabase_interfaces_SelectContext_Call) Return(_a0 error) *MockSqlDatabase_interfaces_SelectContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlDatabase_interfaces_SelectContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockSqlDatabase_interfaces_SelectContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSqlDatabase_interfaces creates a new instance of MockSqlDatabase_interfaces. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSqlDatabase_interfaces(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSqlDatabase_interfaces {
	mock := &MockSqlDatabase_interfaces{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}