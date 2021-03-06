// Code generated by mockery v2.2.1. DO NOT EDIT.

package services

import (
	api_models "hydro_monitor/web_api/pkg/models/api_models"
	db_models "hydro_monitor/web_api/pkg/models/db_models"

	mock "github.com/stretchr/testify/mock"
)

// ReadingsServiceMock is an autogenerated mock type for the ReadingsService type
type ReadingsServiceMock struct {
	mock.Mock
}

// AddPhotoToReading provides a mock function with given fields: nodeId, photoDTO
func (_m *ReadingsServiceMock) AddPhotoToReading(nodeId string, photoDTO *api_models.PhotoDTO) ServiceError {
	ret := _m.Called(nodeId, photoDTO)

	var r0 ServiceError
	if rf, ok := ret.Get(0).(func(string, *api_models.PhotoDTO) ServiceError); ok {
		r0 = rf(nodeId, photoDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ServiceError)
		}
	}

	return r0
}

// CreateReading provides a mock function with given fields: nodeId, reading
func (_m *ReadingsServiceMock) CreateReading(nodeId string, reading *api_models.ReadingDTO) (*api_models.GetReadingDTO, error) {
	ret := _m.Called(nodeId, reading)

	var r0 *api_models.GetReadingDTO
	if rf, ok := ret.Get(0).(func(string, *api_models.ReadingDTO) *api_models.GetReadingDTO); ok {
		r0 = rf(nodeId, reading)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.GetReadingDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *api_models.ReadingDTO) error); ok {
		r1 = rf(nodeId, reading)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteReading provides a mock function with given fields: nodeId, readingId
func (_m *ReadingsServiceMock) DeleteReading(nodeId string, readingId string) ServiceError {
	ret := _m.Called(nodeId, readingId)

	var r0 ServiceError
	if rf, ok := ret.Get(0).(func(string, string) ServiceError); ok {
		r0 = rf(nodeId, readingId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ServiceError)
		}
	}

	return r0
}

// GetNodeReading provides a mock function with given fields: nodeId, readingId
func (_m *ReadingsServiceMock) GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, ServiceError) {
	ret := _m.Called(nodeId, readingId)

	var r0 *api_models.GetReadingDTO
	if rf, ok := ret.Get(0).(func(string, string) *api_models.GetReadingDTO); ok {
		r0 = rf(nodeId, readingId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.GetReadingDTO)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func(string, string) ServiceError); ok {
		r1 = rf(nodeId, readingId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}

// GetNodeReadings provides a mock function with given fields: nodeId, pageState, pageSize
func (_m *ReadingsServiceMock) GetNodeReadings(nodeId string, pageState []byte, pageSize int) (*api_models.PaginatedDTO, ServiceError) {
	ret := _m.Called(nodeId, pageState, pageSize)

	var r0 *api_models.PaginatedDTO
	if rf, ok := ret.Get(0).(func(string, []byte, int) *api_models.PaginatedDTO); ok {
		r0 = rf(nodeId, pageState, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.PaginatedDTO)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func(string, []byte, int) ServiceError); ok {
		r1 = rf(nodeId, pageState, pageSize)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}

// GetNodesLastReading provides a mock function with given fields: nodes
func (_m *ReadingsServiceMock) GetNodesLastReading(nodes []*api_models.NodeDTO) (map[string]*api_models.GetReadingDTO, ServiceError) {
	ret := _m.Called(nodes)

	var r0 map[string]*api_models.GetReadingDTO
	if rf, ok := ret.Get(0).(func([]*api_models.NodeDTO) map[string]*api_models.GetReadingDTO); ok {
		r0 = rf(nodes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*api_models.GetReadingDTO)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func([]*api_models.NodeDTO) ServiceError); ok {
		r1 = rf(nodes)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}

// GetReadingPhoto provides a mock function with given fields: readingId, number
func (_m *ReadingsServiceMock) GetReadingPhoto(readingId string, number int) (*db_models.Photo, ServiceError) {
	ret := _m.Called(readingId, number)

	var r0 *db_models.Photo
	if rf, ok := ret.Get(0).(func(string, int) *db_models.Photo); ok {
		r0 = rf(readingId, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db_models.Photo)
		}
	}

	var r1 ServiceError
	if rf, ok := ret.Get(1).(func(string, int) ServiceError); ok {
		r1 = rf(readingId, number)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(ServiceError)
		}
	}

	return r0, r1
}
