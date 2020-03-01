// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// ServiceError is an autogenerated mock type for the ServiceError type
type ServiceError struct {
	mock.Mock
}

// Error provides a mock function with given fields:
func (_m *ServiceError) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ToHTTPError provides a mock function with given fields:
func (_m *ServiceError) ToHTTPError() *echo.HTTPError {
	ret := _m.Called()

	var r0 *echo.HTTPError
	if rf, ok := ret.Get(0).(func() *echo.HTTPError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.HTTPError)
		}
	}

	return r0
}
