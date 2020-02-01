package services

import (
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
)

type NodeService interface {
	GetNode(nodeId string) (*api_models.NodeDTO, error)
	GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error)
	GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error)
	UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error)
}

type nodeServiceImpl struct {
	nodeRepository repositories.Repository
}

func (n *nodeServiceImpl) GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error) {
	respManualReading := &db_models.ManualReadingDTO{NodeId: nodeId}
	if err := n.nodeRepository.Get(respManualReading); err != nil {
		return nil, err
	}
	return respManualReading.ToAPIManualReadingDTO(), nil
}

func (n *nodeServiceImpl) UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error) {
	node := &db_models.ManualReadingDTO{
		NodeId:        nodeId,
		ManualReading: manualReading,
	}
	if err := n.nodeRepository.Update(node); err != nil {
		return nil, err
	}
	resp := &api_models.ManualReadingDTO{ManualReading: node.ManualReading}
	return resp, nil
}

func NewNodeService(dbClient db.Client) NodeService {
	nodeRepository := repositories.NewNodeRepository(dbClient)
	return &nodeServiceImpl{nodeRepository: nodeRepository}
}

func (n *nodeServiceImpl) GetNode(nodeId string) (*api_models.NodeDTO, error) {
	node := db_models.NodeDTO{Id: nodeId}
	err := n.nodeRepository.Get(&node)
	return node.ToAPINodeDTO(), err
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
