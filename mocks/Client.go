// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	db_models "hydro_monitor/web_api/pkg/models/db_models"

	mock "github.com/stretchr/testify/mock"

	table "github.com/scylladb/gocqlx/table"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Client) Close() {
	_m.Called()
}

// Delete provides a mock function with given fields: _a0, args
func (_m *Client) Delete(_a0 *table.Table, args db_models.DbDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.DbDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0, args
func (_m *Client) Get(_a0 *table.Table, args db_models.DbDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.DbDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: _a0, args
func (_m *Client) Insert(_a0 *table.Table, args db_models.DbDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.DbDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Migrate provides a mock function with given fields: dir
func (_m *Client) Migrate(dir string) {
	_m.Called(dir)
}

// Select provides a mock function with given fields: _a0, args
func (_m *Client) Select(_a0 *table.Table, args db_models.SelectDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.SelectDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: _a0, args
func (_m *Client) SelectAll(_a0 *table.Table, args db_models.SelectDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.SelectDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, args
func (_m *Client) Update(_a0 *table.Table, args db_models.DbDTO) error {
	ret := _m.Called(_a0, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(*table.Table, db_models.DbDTO) error); ok {
		r0 = rf(_a0, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
