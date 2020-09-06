package services

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
	"hydro_monitor/web_api/pkg/utils"
	"net/http"
)

type NodeService interface {
	CreateNodeConfiguration(nodeId string, configuration map[string]*api_models.StateDTO) error
	CreateNode(node *api_models.NodeDTO) (*api_models.NodeDTO, ServiceError)
	DeleteNode(nodeId string) error
	GetNode(nodeId string) (*api_models.NodeDTO, ServiceError)
	GetNodes() ([]*api_models.NodeDTO, error)
	GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error)
	GetNodeConfiguration(nodeId string) (map[string]*api_models.StateDTO, ServiceError)
	UpdateNode(node *api_models.NodeDTO) ServiceError
	UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error)
	CheckNodeCredentials(nodeId string, password string) (bool, ServiceError)
}

type nodeServiceImpl struct {
	nodesRepository          repositories.Repository
	configurationsRepository repositories.Repository
}

func (n *nodeServiceImpl) CheckNodeCredentials(nodeId string, password string) (bool, ServiceError) {
	dbNode := &db_models.NodeDTO{Id: &nodeId}
	if err := n.nodesRepository.Get(dbNode); err != nil {
		if err == gocql.ErrNotFound {
			return false, NewNotFoundError("Node not found", err)
		}
		return false, NewGenericServiceError("Server error when getting node", err)
	}
	if err := bcrypt.CompareHashAndPassword(dbNode.Password, []byte(password)); err != nil {
		return false, NewInvalidCredentialsError("Wrong node id or password", err)
	}
	return true, nil
}

func (n *nodeServiceImpl) UpdateNode(node *api_models.NodeDTO) ServiceError {
	dbNode := &db_models.NodeDTO{
		Id:            node.Id,
		Description:   node.Description,
		ManualReading: node.ManualReading,
	}
	if err := n.nodesRepository.Update(dbNode); err != nil {
		return NewGenericServiceError("Error when trying to update node information", err)
	}
	return nil
}

func (n *nodeServiceImpl) CreateNodeConfiguration(nodeId string, configuration map[string]*api_models.StateDTO) error {
	validator := NewConfigurationValidator()
	if valid := validator.ConfigurationIsValid(configuration); !valid {
		return NewGenericClientError(
			"Configuration is not valid",
			fmt.Errorf("configuration %s is not valid", ConfigToString(configuration)))
	}
	rawConfiguration, err := json.Marshal(configuration)
	if err != nil {
		return NewGenericServiceError("Error when trying to marshal node's configuration", err)
	}
	configurationDTO := &db_models.ConfigurationDTO{NodeId: nodeId, Configuration: string(rawConfiguration)}
	return n.configurationsRepository.Insert(configurationDTO)
}

func (n *nodeServiceImpl) GetNodes() ([]*api_models.NodeDTO, error) {
	nodesDTO := db_models.NewNodesDTO()
	if err := n.nodesRepository.SelectAll(nodesDTO); err != nil {
		return nil, err
	}
	return nodesDTO.ConvertToApiNodes(), nil
}

func (n *nodeServiceImpl) DeleteNode(nodeId string) error {
	dbNode := &db_models.DeleteNodeDTO{Id: nodeId}
	return n.nodesRepository.Delete(dbNode)
}

func (n *nodeServiceImpl) CreateNode(node *api_models.NodeDTO) (*api_models.NodeDTO, ServiceError) {
	dbNode := &db_models.NodeDTO{Id: node.Id}
	if err := n.nodesRepository.Get(dbNode); err == nil {
		return nil, NewServiceError(http.StatusUnprocessableEntity, "Node already exists", nil)
	}
	password := utils.GenerateRandomString(16)
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, NewGenericServiceError("Error when trying to register a new node", err)
	}
	manualReadingFalse := false
	dbNode = &db_models.NodeDTO{
		Id:            node.Id,
		Description:   node.Description,
		ManualReading: &manualReadingFalse,
		Password:      encryptedPassword,
	}
	if err = n.nodesRepository.Insert(dbNode); err != nil {
		return nil, NewGenericServiceError("Error when trying to register a new node", err)
	}
	node.ManualReading = &manualReadingFalse
	node.Password = &password
	return node, nil
}

func (n *nodeServiceImpl) GetNodeManualReadingStatus(nodeId string) (*api_models.ManualReadingDTO, error) {
	respManualReading := &db_models.ManualReadingDTO{NodeId: nodeId}
	if err := n.nodesRepository.Get(respManualReading); err != nil {
		return nil, err
	}
	return respManualReading.ToAPIManualReadingDTO(), nil
}

func (n *nodeServiceImpl) UpdateNodeManualReading(nodeId string, manualReading bool) (*api_models.ManualReadingDTO, error) {
	node := &db_models.ManualReadingDTO{
		NodeId:        nodeId,
		ManualReading: manualReading,
	}
	if err := n.nodesRepository.Update(node); err != nil {
		return nil, err
	}
	resp := &api_models.ManualReadingDTO{ManualReading: node.ManualReading}
	return resp, nil
}

func (n *nodeServiceImpl) GetNode(nodeId string) (*api_models.NodeDTO, ServiceError) {
	node := db_models.NodeDTO{Id: &nodeId}
	err := n.nodesRepository.Get(&node)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node not found", err)
		}
		return nil, NewGenericServiceError("Server error when getting node", err)
	}
	return node.ToSingleAPINodeDTO(), nil
}

func (n *nodeServiceImpl) GetNodeConfiguration(nodeId string) (map[string]*api_models.StateDTO, ServiceError) {
	configuration := make(map[string]*api_models.StateDTO)
	configurationDTO := &db_models.ConfigurationDTO{NodeId: nodeId}
	if err := n.configurationsRepository.Get(configurationDTO); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node configuration not found", gocql.ErrNotFound)
		}
		return nil, NewGenericServiceError("Server error when getting node configuration", err)
	}
	if err := json.Unmarshal([]byte(configurationDTO.Configuration), &configuration); err != nil {
		return nil, NewGenericServiceError("Server error when unmarshaling node configuration", err)
	}
	return configuration, nil
}

func NewNodeService(configurationsRepository repositories.Repository, nodesRepository repositories.Repository) NodeService {
	return &nodeServiceImpl{
		nodesRepository:          nodesRepository,
		configurationsRepository: configurationsRepository,
	}
}
