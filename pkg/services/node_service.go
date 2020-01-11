package services

import (
	db_client "hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models"
	"hydro_monitor/web_api/pkg/repositories"
)

type NodeService interface {
	GetNode(nodeId string) (*models.Node, error)
	GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error)
}

type nodeServiceImpl struct {
	nodeRepository repositories.Repository
}

func NewNodeService(dbClient db_client.DbClient) NodeService {
	nodeRepository := repositories.NewNodeRepository(dbClient)
	return &nodeServiceImpl{nodeRepository: nodeRepository}
}

func (n *nodeServiceImpl) GetNode(nodeId string) (*models.Node, error) {
	node := models.Node{Id: nodeId}
	err := n.nodeRepository.Get(node)
	return &node, err
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
