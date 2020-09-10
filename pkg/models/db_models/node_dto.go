package db_models

import "hydro_monitor/web_api/pkg/models/api_models"

type NodeDTO struct {
	Id            *string `db:"id"`
	Description   *string `db:"description"`
	ManualReading *bool   `db:"manual_reading"`
	Password      []byte  `db:"password"`
}

func (n *NodeDTO) GetColumns() []string {
	nodeColumns := make([]string, 0)
	if n.Description != nil {
		nodeColumns = append(nodeColumns, "description")
	}
	if n.ManualReading != nil {
		nodeColumns = append(nodeColumns, "manual_reading")
	}
	if n.Password != nil {
		nodeColumns = append(nodeColumns, "password")
	}
	return nodeColumns
}

func (n *NodeDTO) ToSingleAPINodeDTO() *api_models.NodeDTO {
	return &api_models.NodeDTO{
		Description:   n.Description,
		ManualReading: n.ManualReading,
	}
}

func (n *NodeDTO) ToAPINodeDTO() *api_models.NodeDTO {
	return &api_models.NodeDTO{
		Id:            n.Id,
		Description:   n.Description,
		ManualReading: n.ManualReading,
	}
}
