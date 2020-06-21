// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	api_models "hydro_monitor/web_api/pkg/models/api_models"

	mock "github.com/stretchr/testify/mock"

	services "hydro_monitor/web_api/pkg/services"
)

// UsersService is an autogenerated mock type for the UsersService type
type UsersService struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: email
func (_m *UsersService) DeleteUser(email string) services.ServiceError {
	ret := _m.Called(email)

	var r0 services.ServiceError
	if rf, ok := ret.Get(0).(func(string) services.ServiceError); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.ServiceError)
		}
	}

	return r0
}

// GetUserInfo provides a mock function with given fields: email
func (_m *UsersService) GetUserInfo(email string) (*api_models.UserDTO, services.ServiceError) {
	ret := _m.Called(email)

	var r0 *api_models.UserDTO
	if rf, ok := ret.Get(0).(func(string) *api_models.UserDTO); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.UserDTO)
		}
	}

	var r1 services.ServiceError
	if rf, ok := ret.Get(1).(func(string) services.ServiceError); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.ServiceError)
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: user
func (_m *UsersService) Register(user *api_models.UserDTO) services.ServiceError {
	ret := _m.Called(user)

	var r0 services.ServiceError
	if rf, ok := ret.Get(0).(func(*api_models.UserDTO) services.ServiceError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.ServiceError)
		}
	}

	return r0
}

// UpdateUser provides a mock function with given fields: user
func (_m *UsersService) UpdateUser(user *api_models.UserDTO) services.ServiceError {
	ret := _m.Called(user)

	var r0 services.ServiceError
	if rf, ok := ret.Get(0).(func(*api_models.UserDTO) services.ServiceError); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.ServiceError)
		}
	}

	return r0
}

// ValidateCredentials provides a mock function with given fields: email, password
func (_m *UsersService) ValidateCredentials(email string, password string) (*api_models.UserDTO, services.ServiceError) {
	ret := _m.Called(email, password)

	var r0 *api_models.UserDTO
	if rf, ok := ret.Get(0).(func(string, string) *api_models.UserDTO); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.UserDTO)
		}
	}

	var r1 services.ServiceError
	if rf, ok := ret.Get(1).(func(string, string) services.ServiceError); ok {
		r1 = rf(email, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(services.ServiceError)
		}
	}

	return r0, r1
}
