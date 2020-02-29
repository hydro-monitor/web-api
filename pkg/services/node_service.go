package services

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
)

type NodeService interface {
	CreateNodeConfiguration(states []*api_models.State) error
	CreateNode(node *api_models.NodeDTO) error
	DeleteNode(nodeId string) error
	GetNode(nodeId string) (*api_models.NodeDTO, ServiceError)
	GetNodes() ([]*api_models.NodeDTO, error)
	GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error)
	GetNodeConfiguration(nodeId string) ([]*api_models.State, error)
	UpdateNodeConfiguration(states []*api_models.State) error
	UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error)
}

type nodeServiceImpl struct {
	nodeRepository   repositories.Repository
	statesRepository repositories.Repository
}

func (n *nodeServiceImpl) UpdateNodeConfiguration(states []*api_models.State) error {
	for _, state := range states {
		dbState := &db_models.StateDTO{
			NodeId:           state.NodeId,
			Name:             state.Name,
			PhotosPerReading: state.PicturesNum,
			ReadingInterval:  state.Interval,
			LowerLimit:       state.LowerLimit,
			UpperLimit:       state.UpperLimit,
			NextState:        state.Next,
			PreviousState:    state.Prev,
		}
		if err := n.statesRepository.Update(dbState); err != nil {
			return err
		}
	}
	return nil
}

func (n *nodeServiceImpl) CreateNodeConfiguration(states []*api_models.State) error {
	for _, state := range states {
		dbState := &db_models.StateDTO{
			NodeId:           state.NodeId,
			Name:             state.Name,
			PhotosPerReading: state.PicturesNum,
			ReadingInterval:  state.Interval,
			LowerLimit:       state.LowerLimit,
			UpperLimit:       state.UpperLimit,
			NextState:        state.Next,
			PreviousState:    state.Prev,
		}
		if err := n.statesRepository.Insert(dbState); err != nil {
			return err
		}
	}
	return nil
}

func (n *nodeServiceImpl) GetNodes() ([]*api_models.NodeDTO, error) {
	nodesDTO := db_models.NewNodesDTO()
	if err := n.nodeRepository.SelectAll(nodesDTO); err != nil {
		return nil, err
	}
	return nodesDTO.ConvertToApiNodes(), nil
}

func (n *nodeServiceImpl) DeleteNode(nodeId string) error {
	dbNode := &db_models.NodeDTO{Id: nodeId}
	return n.nodeRepository.Delete(dbNode)
}

func (n *nodeServiceImpl) CreateNode(node *api_models.NodeDTO) error {
	dbNode := &db_models.NodeDTO{
		Id:            node.Id,
		Description:   node.Description,
		ManualReading: false,
	}
	return n.nodeRepository.Insert(dbNode)
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
	statesRepository := repositories.NewStatesRepository(dbClient)
	return &nodeServiceImpl{nodeRepository: nodeRepository, statesRepository: statesRepository}
}

func (n *nodeServiceImpl) GetNode(nodeId string) (*api_models.NodeDTO, ServiceError) {
	node := db_models.NodeDTO{Id: nodeId}
	err := n.nodeRepository.Get(&node)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node not found", err)
		} else {
			return nil, NewGenericServiceError("Server error when getting node", err)
		}
	}
	return node.ToAPINodeDTO(), nil
}

func (n *nodeServiceImpl) GetNodeConfiguration(nodeId string) ([]*api_models.State, error) {
	statesDto := db_models.NewStatesDTO(nodeId)
	if err := n.statesRepository.Select(statesDto); err != nil {
		return nil, err
	}
	return statesDto.ConvertToAPIStates(), nil
}
