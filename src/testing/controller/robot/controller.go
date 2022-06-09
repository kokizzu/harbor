// Code generated by mockery v2.12.3. DO NOT EDIT.

package robot

import (
	context "context"

	q "github.com/goharbor/harbor/src/lib/q"
	mock "github.com/stretchr/testify/mock"

	robot "github.com/goharbor/harbor/src/controller/robot"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Controller) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, r
func (_m *Controller) Create(ctx context.Context, r *robot.Robot) (int64, string, error) {
	ret := _m.Called(ctx, r)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *robot.Robot) int64); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *robot.Robot) string); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *robot.Robot) error); ok {
		r2 = rf(ctx, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id, option
func (_m *Controller) Get(ctx context.Context, id int64, option *robot.Option) (*robot.Robot, error) {
	ret := _m.Called(ctx, id, option)

	var r0 *robot.Robot
	if rf, ok := ret.Get(0).(func(context.Context, int64, *robot.Option) *robot.Robot); ok {
		r0 = rf(ctx, id, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.Robot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, *robot.Option) error); ok {
		r1 = rf(ctx, id, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query, option
func (_m *Controller) List(ctx context.Context, query *q.Query, option *robot.Option) ([]*robot.Robot, error) {
	ret := _m.Called(ctx, query, option)

	var r0 []*robot.Robot
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query, *robot.Option) []*robot.Robot); ok {
		r0 = rf(ctx, query, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*robot.Robot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query, *robot.Option) error); ok {
		r1 = rf(ctx, query, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, r, option
func (_m *Controller) Update(ctx context.Context, r *robot.Robot, option *robot.Option) error {
	ret := _m.Called(ctx, r, option)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *robot.Robot, *robot.Option) error); ok {
		r0 = rf(ctx, r, option)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewControllerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewController(t NewControllerT) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
