// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	taskfile "github.com/iamlongalong/task/v3/taskfile"
)

// StatusCheckable is an autogenerated mock type for the StatusCheckable type
type StatusCheckable struct {
	mock.Mock
}

type StatusCheckable_Expecter struct {
	mock *mock.Mock
}

func (_m *StatusCheckable) EXPECT() *StatusCheckable_Expecter {
	return &StatusCheckable_Expecter{mock: &_m.Mock}
}

// IsUpToDate provides a mock function with given fields: ctx, t
func (_m *StatusCheckable) IsUpToDate(ctx context.Context, t *taskfile.Task) (bool, error) {
	ret := _m.Called(ctx, t)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *taskfile.Task) (bool, error)); ok {
		return rf(ctx, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *taskfile.Task) bool); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *taskfile.Task) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusCheckable_IsUpToDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUpToDate'
type StatusCheckable_IsUpToDate_Call struct {
	*mock.Call
}

// IsUpToDate is a helper method to define mock.On call
//   - ctx context.Context
//   - t *taskfile.Task
func (_e *StatusCheckable_Expecter) IsUpToDate(ctx interface{}, t interface{}) *StatusCheckable_IsUpToDate_Call {
	return &StatusCheckable_IsUpToDate_Call{Call: _e.mock.On("IsUpToDate", ctx, t)}
}

func (_c *StatusCheckable_IsUpToDate_Call) Run(run func(ctx context.Context, t *taskfile.Task)) *StatusCheckable_IsUpToDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*taskfile.Task))
	})
	return _c
}

func (_c *StatusCheckable_IsUpToDate_Call) Return(_a0 bool, _a1 error) *StatusCheckable_IsUpToDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatusCheckable_IsUpToDate_Call) RunAndReturn(run func(context.Context, *taskfile.Task) (bool, error)) *StatusCheckable_IsUpToDate_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStatusCheckable interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatusCheckable creates a new instance of StatusCheckable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatusCheckable(t mockConstructorTestingTNewStatusCheckable) *StatusCheckable {
	mock := &StatusCheckable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
