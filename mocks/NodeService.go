// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	api_models "hydro_monitor/web_api/pkg/models/api_models"

	mock "github.com/stretchr/testify/mock"
)

// NodeService is an autogenerated mock type for the NodeService type
type NodeService struct {
	mock.Mock
}

// GetNode provides a mock function with given fields: nodeId
func (_m *NodeService) GetNode(nodeId string) (*api_models.NodeDTO, error) {
	ret := _m.Called(nodeId)

	var r0 *api_models.NodeDTO
	if rf, ok := ret.Get(0).(func(string) *api_models.NodeDTO); ok {
		r0 = rf(nodeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.NodeDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nodeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeConfiguration provides a mock function with given fields: nodeId
func (_m *NodeService) GetNodeConfiguration(nodeId string) ([]*api_models.State, error) {
	ret := _m.Called(nodeId)

	var r0 []*api_models.State
	if rf, ok := ret.Get(0).(func(string) []*api_models.State); ok {
		r0 = rf(nodeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api_models.State)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nodeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeManualReadingStatus provides a mock function with given fields: nodeId
func (_m *NodeService) GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error) {
	ret := _m.Called(nodeId)

	var r0 *api_models.ManualReadingDTO
	if rf, ok := ret.Get(0).(func(string) *api_models.ManualReadingDTO); ok {
		r0 = rf(nodeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.ManualReadingDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nodeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNodeManualReading provides a mock function with given fields: nodeId, manualReading
func (_m *NodeService) UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error) {
	ret := _m.Called(nodeId, manualReading)

	var r0 *api_models.ManualReadingDTO
	if rf, ok := ret.Get(0).(func(string, bool) *api_models.ManualReadingDTO); ok {
		r0 = rf(nodeId, manualReading)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api_models.ManualReadingDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(nodeId, manualReading)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
