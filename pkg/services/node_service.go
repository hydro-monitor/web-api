package services

import "hydro_monitor/web_api/pkg/models"

type NodeService interface {
	GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error)
}

type nodeServiceImpl struct {
}

func NewNodeService() NodeService {
	return &nodeServiceImpl{}
}

func (n *nodeServiceImpl) GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error) {
	state1 := models.State{
		Name:        "Normal",
		Interval:    60,
		UpperLimit:  9999999,
		LowerLimit:  -9999999,
		PicturesNum: 0,
		Next:        "Alto",
		Prev:        "Bajo",
	}
	statesMap := make(map[string]*models.State)
	statesMap[state1.Name] = &state1
	expectedNodeConfiguration := models.NodeConfiguration{
		NodeId: "1",
		States: statesMap,
	}
	return &expectedNodeConfiguration, nil
}
