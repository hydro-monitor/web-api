// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	db_models "hydro_monitor/web_api/pkg/models/db_models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: args
func (_m *Repository) Get(args interface{}) error {
	ret := _m.Called(args)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: args
func (_m *Repository) Insert(args db_models.DbDTO) error {
	ret := _m.Called(args)

	var r0 error
	if rf, ok := ret.Get(0).(func(db_models.DbDTO) error); ok {
		r0 = rf(args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: args
func (_m *Repository) Update(args db_models.DbDTO) error {
	ret := _m.Called(args)

	var r0 error
	if rf, ok := ret.Get(0).(func(db_models.DbDTO) error); ok {
		r0 = rf(args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
