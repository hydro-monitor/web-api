package services

import "hydro_monitor/web_api/pkg/models"

type NodeService interface {
	GetNode(nodeId string) (*models.Node, error)
	GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error)
}

type nodeServiceImpl struct {
}

func NewNodeService() NodeService {
	return &nodeServiceImpl{}
}

func (n *nodeServiceImpl) GetNode(nodeId string) (*models.Node, error) {
	node := models.Node{
		Id:            nodeId,
		Description:   "A node",
		Configuration: "1",
		State:         "Normal",
	}
	return &node, nil
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
		NodeId: nodeId,
		States: statesMap,
	}
	return &expectedNodeConfiguration, nil
}
