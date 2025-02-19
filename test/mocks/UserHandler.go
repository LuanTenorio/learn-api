// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// UserHandler is an autogenerated mock type for the UserHandler type
type UserHandler struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: c
func (_m *UserHandler) CreateUser(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserHandler creates a new instance of UserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserHandler {
	mock := &UserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
