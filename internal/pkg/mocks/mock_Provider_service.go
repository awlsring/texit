// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	provider "github.com/awlsring/texit/internal/pkg/domain/provider"
	mock "github.com/stretchr/testify/mock"
)

// MockProvider_service is an autogenerated mock type for the Provider type
type MockProvider_service struct {
	mock.Mock
}

type MockProvider_service_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProvider_service) EXPECT() *MockProvider_service_Expecter {
	return &MockProvider_service_Expecter{mock: &_m.Mock}
}

// Describe provides a mock function with given fields: _a0, _a1
func (_m *MockProvider_service) Describe(_a0 context.Context, _a1 provider.Identifier) (*provider.Provider, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Describe")
	}

	var r0 *provider.Provider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provider.Identifier) (*provider.Provider, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provider.Identifier) *provider.Provider); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Provider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provider.Identifier) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProvider_service_Describe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Describe'
type MockProvider_service_Describe_Call struct {
	*mock.Call
}

// Describe is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 provider.Identifier
func (_e *MockProvider_service_Expecter) Describe(_a0 interface{}, _a1 interface{}) *MockProvider_service_Describe_Call {
	return &MockProvider_service_Describe_Call{Call: _e.mock.On("Describe", _a0, _a1)}
}

func (_c *MockProvider_service_Describe_Call) Run(run func(_a0 context.Context, _a1 provider.Identifier)) *MockProvider_service_Describe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provider.Identifier))
	})
	return _c
}

func (_c *MockProvider_service_Describe_Call) Return(_a0 *provider.Provider, _a1 error) *MockProvider_service_Describe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProvider_service_Describe_Call) RunAndReturn(run func(context.Context, provider.Identifier) (*provider.Provider, error)) *MockProvider_service_Describe_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0
func (_m *MockProvider_service) List(_a0 context.Context) ([]*provider.Provider, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*provider.Provider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*provider.Provider, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*provider.Provider); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*provider.Provider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProvider_service_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockProvider_service_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockProvider_service_Expecter) List(_a0 interface{}) *MockProvider_service_List_Call {
	return &MockProvider_service_List_Call{Call: _e.mock.On("List", _a0)}
}

func (_c *MockProvider_service_List_Call) Run(run func(_a0 context.Context)) *MockProvider_service_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockProvider_service_List_Call) Return(_a0 []*provider.Provider, _a1 error) *MockProvider_service_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProvider_service_List_Call) RunAndReturn(run func(context.Context) ([]*provider.Provider, error)) *MockProvider_service_List_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProvider_service creates a new instance of MockProvider_service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProvider_service(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProvider_service {
	mock := &MockProvider_service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
