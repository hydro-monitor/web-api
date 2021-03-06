// Code generated by mockery v2.2.1. DO NOT EDIT.

package services

import (
	api_models "hydro_monitor/web_api/pkg/models/api_models"

	mock "github.com/stretchr/testify/mock"
)

// UsersServiceMock is an autogenerated mock type for the UsersService type
type UsersServiceMock struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: email
func (_m *UsersServiceMock) DeleteUser(email string) ServiceError {
	ret := _m.Called(email)

	var r0 ServiceError
	if rf, ok := ret.Get(0).(func(string) ServiceError); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ServiceError)
		}
	}

	return r0
}

// GetUserInfo provides a mock function with given fields: email
func (_m *UsersServiceMock) GetUserInfo(email string) (*api_models.UserDTO, ServiceError) {
	ret := _m.Called(email)

	var r0 *api_models.UserDTO
	if rf, ok := ret.Get(0).(func(string) *api_models.UserDTO); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.UserDTO)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func(string) ServiceError); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *UsersServiceMock) Register(user *api_models.UserDTO) ServiceError {
	ret := _m.Called(user)

	var r0 ServiceError
	if rf, ok := ret.Get(0).(func(*api_models.UserDTO) ServiceError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ServiceError)
		}
	}

	return r0
}

// UpdateUser provides a mock function with given fields: user
func (_m *UsersServiceMock) UpdateUser(user *api_models.UserDTO) ServiceError {
	ret := _m.Called(user)

	var r0 ServiceError
	if rf, ok := ret.Get(0).(func(*api_models.UserDTO) ServiceError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ServiceError)
		}
	}

	return r0
}

// ValidateCredentials provides a mock function with given fields: email, password
func (_m *UsersServiceMock) ValidateCredentials(email string, password string) (*api_models.UserDTO, ServiceError) {
	ret := _m.Called(email, password)

	var r0 *api_models.UserDTO
	if rf, ok := ret.Get(0).(func(string, string) *api_models.UserDTO); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.UserDTO)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func(string, string) ServiceError); ok {
		r1 = rf(email, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}
