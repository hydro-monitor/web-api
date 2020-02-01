// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	api_models "hydro_monitor/web_api/pkg/models/api_models"
	db_models "hydro_monitor/web_api/pkg/models/db_models"

	mock "github.com/stretchr/testify/mock"
)

// ReadingsService is an autogenerated mock type for the ReadingsService type
type ReadingsService struct {
	mock.Mock
}

// AddPhotoToReading provides a mock function with given fields: photoDTO
func (_m *ReadingsService) AddPhotoToReading(photoDTO *api_models.PhotoDTO) (*api_models.PhotoMetadataDTO, error) {
	ret := _m.Called(photoDTO)

	var r0 *api_models.PhotoMetadataDTO
	if rf, ok := ret.Get(0).(func(*api_models.PhotoDTO) *api_models.PhotoMetadataDTO); ok {
		r0 = rf(photoDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.PhotoMetadataDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*api_models.PhotoDTO) error); ok {
		r1 = rf(photoDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReading provides a mock function with given fields: nodeId, reading
func (_m *ReadingsService) CreateReading(nodeId string, reading *api_models.Reading) (*api_models.GetReadingDTO, error) {
	ret := _m.Called(nodeId, reading)

	var r0 *api_models.GetReadingDTO
	if rf, ok := ret.Get(0).(func(string, *api_models.Reading) *api_models.GetReadingDTO); ok {
		r0 = rf(nodeId, reading)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.GetReadingDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *api_models.Reading) error); ok {
		r1 = rf(nodeId, reading)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeReading provides a mock function with given fields: nodeId, readingId
func (_m *ReadingsService) GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, error) {
	ret := _m.Called(nodeId, readingId)

	var r0 *api_models.GetReadingDTO
	if rf, ok := ret.Get(0).(func(string, string) *api_models.GetReadingDTO); ok {
		r0 = rf(nodeId, readingId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.GetReadingDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(nodeId, readingId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeReadings provides a mock function with given fields: nodeId
func (_m *ReadingsService) GetNodeReadings(nodeId string) error {
	ret := _m.Called(nodeId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(nodeId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReadingPhoto provides a mock function with given fields: readingId, number
func (_m *ReadingsService) GetReadingPhoto(readingId string, number int) (*db_models.Photo, error) {
	ret := _m.Called(readingId, number)

	var r0 *db_models.Photo
	if rf, ok := ret.Get(0).(func(string, int) *db_models.Photo); ok {
		r0 = rf(readingId, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*db_models.Photo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(readingId, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}