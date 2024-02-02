// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	iam "github.com/aws/aws-sdk-go-v2/service/iam"

	mock "github.com/stretchr/testify/mock"
)

// MockIamClient_interfaces is an autogenerated mock type for the IamClient type
type MockIamClient_interfaces struct {
	mock.Mock
}

type MockIamClient_interfaces_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIamClient_interfaces) EXPECT() *MockIamClient_interfaces_Expecter {
	return &MockIamClient_interfaces_Expecter{mock: &_m.Mock}
}

// AttachRolePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) AttachRolePolicy(ctx context.Context, params *iam.AttachRolePolicyInput, optFns ...func(*iam.Options)) (*iam.AttachRolePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AttachRolePolicy")
	}

	var r0 *iam.AttachRolePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.AttachRolePolicyInput, ...func(*iam.Options)) (*iam.AttachRolePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.AttachRolePolicyInput, ...func(*iam.Options)) *iam.AttachRolePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.AttachRolePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.AttachRolePolicyInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_AttachRolePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AttachRolePolicy'
type MockIamClient_interfaces_AttachRolePolicy_Call struct {
	*mock.Call
}

// AttachRolePolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.AttachRolePolicyInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) AttachRolePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_AttachRolePolicy_Call {
	return &MockIamClient_interfaces_AttachRolePolicy_Call{Call: _e.mock.On("AttachRolePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_AttachRolePolicy_Call) Run(run func(ctx context.Context, params *iam.AttachRolePolicyInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_AttachRolePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.AttachRolePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_AttachRolePolicy_Call) Return(_a0 *iam.AttachRolePolicyOutput, _a1 error) *MockIamClient_interfaces_AttachRolePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_AttachRolePolicy_Call) RunAndReturn(run func(context.Context, *iam.AttachRolePolicyInput, ...func(*iam.Options)) (*iam.AttachRolePolicyOutput, error)) *MockIamClient_interfaces_AttachRolePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) CreatePolicy(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options)) (*iam.CreatePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePolicy")
	}

	var r0 *iam.CreatePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreatePolicyInput, ...func(*iam.Options)) (*iam.CreatePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreatePolicyInput, ...func(*iam.Options)) *iam.CreatePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.CreatePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.CreatePolicyInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_CreatePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePolicy'
type MockIamClient_interfaces_CreatePolicy_Call struct {
	*mock.Call
}

// CreatePolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.CreatePolicyInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) CreatePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_CreatePolicy_Call {
	return &MockIamClient_interfaces_CreatePolicy_Call{Call: _e.mock.On("CreatePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_CreatePolicy_Call) Run(run func(ctx context.Context, params *iam.CreatePolicyInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_CreatePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.CreatePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_CreatePolicy_Call) Return(_a0 *iam.CreatePolicyOutput, _a1 error) *MockIamClient_interfaces_CreatePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_CreatePolicy_Call) RunAndReturn(run func(context.Context, *iam.CreatePolicyInput, ...func(*iam.Options)) (*iam.CreatePolicyOutput, error)) *MockIamClient_interfaces_CreatePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRole provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) CreateRole(ctx context.Context, params *iam.CreateRoleInput, optFns ...func(*iam.Options)) (*iam.CreateRoleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateRole")
	}

	var r0 *iam.CreateRoleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateRoleInput, ...func(*iam.Options)) (*iam.CreateRoleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.CreateRoleInput, ...func(*iam.Options)) *iam.CreateRoleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.CreateRoleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.CreateRoleInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_CreateRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRole'
type MockIamClient_interfaces_CreateRole_Call struct {
	*mock.Call
}

// CreateRole is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.CreateRoleInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) CreateRole(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_CreateRole_Call {
	return &MockIamClient_interfaces_CreateRole_Call{Call: _e.mock.On("CreateRole",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_CreateRole_Call) Run(run func(ctx context.Context, params *iam.CreateRoleInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_CreateRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.CreateRoleInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_CreateRole_Call) Return(_a0 *iam.CreateRoleOutput, _a1 error) *MockIamClient_interfaces_CreateRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_CreateRole_Call) RunAndReturn(run func(context.Context, *iam.CreateRoleInput, ...func(*iam.Options)) (*iam.CreateRoleOutput, error)) *MockIamClient_interfaces_CreateRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) DeletePolicy(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options)) (*iam.DeletePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePolicy")
	}

	var r0 *iam.DeletePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeletePolicyInput, ...func(*iam.Options)) (*iam.DeletePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeletePolicyInput, ...func(*iam.Options)) *iam.DeletePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.DeletePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.DeletePolicyInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_DeletePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePolicy'
type MockIamClient_interfaces_DeletePolicy_Call struct {
	*mock.Call
}

// DeletePolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.DeletePolicyInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) DeletePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_DeletePolicy_Call {
	return &MockIamClient_interfaces_DeletePolicy_Call{Call: _e.mock.On("DeletePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_DeletePolicy_Call) Run(run func(ctx context.Context, params *iam.DeletePolicyInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_DeletePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.DeletePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_DeletePolicy_Call) Return(_a0 *iam.DeletePolicyOutput, _a1 error) *MockIamClient_interfaces_DeletePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_DeletePolicy_Call) RunAndReturn(run func(context.Context, *iam.DeletePolicyInput, ...func(*iam.Options)) (*iam.DeletePolicyOutput, error)) *MockIamClient_interfaces_DeletePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRole provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) DeleteRole(ctx context.Context, params *iam.DeleteRoleInput, optFns ...func(*iam.Options)) (*iam.DeleteRoleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRole")
	}

	var r0 *iam.DeleteRoleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteRoleInput, ...func(*iam.Options)) (*iam.DeleteRoleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteRoleInput, ...func(*iam.Options)) *iam.DeleteRoleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.DeleteRoleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.DeleteRoleInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_DeleteRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRole'
type MockIamClient_interfaces_DeleteRole_Call struct {
	*mock.Call
}

// DeleteRole is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.DeleteRoleInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) DeleteRole(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_DeleteRole_Call {
	return &MockIamClient_interfaces_DeleteRole_Call{Call: _e.mock.On("DeleteRole",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_DeleteRole_Call) Run(run func(ctx context.Context, params *iam.DeleteRoleInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_DeleteRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.DeleteRoleInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_DeleteRole_Call) Return(_a0 *iam.DeleteRoleOutput, _a1 error) *MockIamClient_interfaces_DeleteRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_DeleteRole_Call) RunAndReturn(run func(context.Context, *iam.DeleteRoleInput, ...func(*iam.Options)) (*iam.DeleteRoleOutput, error)) *MockIamClient_interfaces_DeleteRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRolePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) DeleteRolePolicy(ctx context.Context, params *iam.DeleteRolePolicyInput, optFns ...func(*iam.Options)) (*iam.DeleteRolePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRolePolicy")
	}

	var r0 *iam.DeleteRolePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteRolePolicyInput, ...func(*iam.Options)) (*iam.DeleteRolePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DeleteRolePolicyInput, ...func(*iam.Options)) *iam.DeleteRolePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.DeleteRolePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.DeleteRolePolicyInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_DeleteRolePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRolePolicy'
type MockIamClient_interfaces_DeleteRolePolicy_Call struct {
	*mock.Call
}

// DeleteRolePolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.DeleteRolePolicyInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) DeleteRolePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_DeleteRolePolicy_Call {
	return &MockIamClient_interfaces_DeleteRolePolicy_Call{Call: _e.mock.On("DeleteRolePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_DeleteRolePolicy_Call) Run(run func(ctx context.Context, params *iam.DeleteRolePolicyInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_DeleteRolePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.DeleteRolePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_DeleteRolePolicy_Call) Return(_a0 *iam.DeleteRolePolicyOutput, _a1 error) *MockIamClient_interfaces_DeleteRolePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_DeleteRolePolicy_Call) RunAndReturn(run func(context.Context, *iam.DeleteRolePolicyInput, ...func(*iam.Options)) (*iam.DeleteRolePolicyOutput, error)) *MockIamClient_interfaces_DeleteRolePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// DetachRolePolicy provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) DetachRolePolicy(ctx context.Context, params *iam.DetachRolePolicyInput, optFns ...func(*iam.Options)) (*iam.DetachRolePolicyOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DetachRolePolicy")
	}

	var r0 *iam.DetachRolePolicyOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DetachRolePolicyInput, ...func(*iam.Options)) (*iam.DetachRolePolicyOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.DetachRolePolicyInput, ...func(*iam.Options)) *iam.DetachRolePolicyOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.DetachRolePolicyOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.DetachRolePolicyInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_DetachRolePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetachRolePolicy'
type MockIamClient_interfaces_DetachRolePolicy_Call struct {
	*mock.Call
}

// DetachRolePolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.DetachRolePolicyInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) DetachRolePolicy(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_DetachRolePolicy_Call {
	return &MockIamClient_interfaces_DetachRolePolicy_Call{Call: _e.mock.On("DetachRolePolicy",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_DetachRolePolicy_Call) Run(run func(ctx context.Context, params *iam.DetachRolePolicyInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_DetachRolePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.DetachRolePolicyInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_DetachRolePolicy_Call) Return(_a0 *iam.DetachRolePolicyOutput, _a1 error) *MockIamClient_interfaces_DetachRolePolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_DetachRolePolicy_Call) RunAndReturn(run func(context.Context, *iam.DetachRolePolicyInput, ...func(*iam.Options)) (*iam.DetachRolePolicyOutput, error)) *MockIamClient_interfaces_DetachRolePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// GetRole provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) GetRole(ctx context.Context, params *iam.GetRoleInput, optFns ...func(*iam.Options)) (*iam.GetRoleOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRole")
	}

	var r0 *iam.GetRoleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.GetRoleInput, ...func(*iam.Options)) (*iam.GetRoleOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.GetRoleInput, ...func(*iam.Options)) *iam.GetRoleOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.GetRoleOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.GetRoleInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_GetRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRole'
type MockIamClient_interfaces_GetRole_Call struct {
	*mock.Call
}

// GetRole is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.GetRoleInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) GetRole(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_GetRole_Call {
	return &MockIamClient_interfaces_GetRole_Call{Call: _e.mock.On("GetRole",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_GetRole_Call) Run(run func(ctx context.Context, params *iam.GetRoleInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_GetRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.GetRoleInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_GetRole_Call) Return(_a0 *iam.GetRoleOutput, _a1 error) *MockIamClient_interfaces_GetRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_GetRole_Call) RunAndReturn(run func(context.Context, *iam.GetRoleInput, ...func(*iam.Options)) (*iam.GetRoleOutput, error)) *MockIamClient_interfaces_GetRole_Call {
	_c.Call.Return(run)
	return _c
}

// ListAttachedRolePolicies provides a mock function with given fields: ctx, params, optFns
func (_m *MockIamClient_interfaces) ListAttachedRolePolicies(ctx context.Context, params *iam.ListAttachedRolePoliciesInput, optFns ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAttachedRolePolicies")
	}

	var r0 *iam.ListAttachedRolePoliciesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListAttachedRolePoliciesInput, ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ListAttachedRolePoliciesInput, ...func(*iam.Options)) *iam.ListAttachedRolePoliciesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.ListAttachedRolePoliciesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *iam.ListAttachedRolePoliciesInput, ...func(*iam.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIamClient_interfaces_ListAttachedRolePolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAttachedRolePolicies'
type MockIamClient_interfaces_ListAttachedRolePolicies_Call struct {
	*mock.Call
}

// ListAttachedRolePolicies is a helper method to define mock.On call
//   - ctx context.Context
//   - params *iam.ListAttachedRolePoliciesInput
//   - optFns ...func(*iam.Options)
func (_e *MockIamClient_interfaces_Expecter) ListAttachedRolePolicies(ctx interface{}, params interface{}, optFns ...interface{}) *MockIamClient_interfaces_ListAttachedRolePolicies_Call {
	return &MockIamClient_interfaces_ListAttachedRolePolicies_Call{Call: _e.mock.On("ListAttachedRolePolicies",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *MockIamClient_interfaces_ListAttachedRolePolicies_Call) Run(run func(ctx context.Context, params *iam.ListAttachedRolePoliciesInput, optFns ...func(*iam.Options))) *MockIamClient_interfaces_ListAttachedRolePolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*iam.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*iam.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*iam.ListAttachedRolePoliciesInput), variadicArgs...)
	})
	return _c
}

func (_c *MockIamClient_interfaces_ListAttachedRolePolicies_Call) Return(_a0 *iam.ListAttachedRolePoliciesOutput, _a1 error) *MockIamClient_interfaces_ListAttachedRolePolicies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIamClient_interfaces_ListAttachedRolePolicies_Call) RunAndReturn(run func(context.Context, *iam.ListAttachedRolePoliciesInput, ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error)) *MockIamClient_interfaces_ListAttachedRolePolicies_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIamClient_interfaces creates a new instance of MockIamClient_interfaces. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIamClient_interfaces(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIamClient_interfaces {
	mock := &MockIamClient_interfaces{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}