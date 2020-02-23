package db_models

import (
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/models/api_models"
)

type NodesDTO struct {
	nodes []*NodeDTO
}

func (n *NodesDTO) GetColumns() []string {
	return nil
}

func (n *NodesDTO) GetBindMap() qb.M {
	return nil
}

func (n *NodesDTO) GetArgs() interface{} {
	return &n.nodes
}

func (n *NodesDTO) ConvertToApiNodes() []*api_models.NodeDTO {
	var nodes []*api_models.NodeDTO
	for _, node := range n.nodes {
		nodes = append(nodes, node.ToAPINodeDTO())
	}
	return nodes
}

func NewNodesDTO() *NodesDTO {
	nodes := make([]*NodeDTO, 0)
	return &NodesDTO{nodes:nodes}
}