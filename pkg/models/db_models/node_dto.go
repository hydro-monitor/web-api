package db_models

import "hydro_monitor/web_api/pkg/models/api_models"

var nodeColumns = []string{"description", "manual_reading"}

type NodeDTO struct {
	Id            string `db:"id"`
	Description   string `db:"description"`
	ManualReading bool   `db:"manual_reading"`
}

func (n *NodeDTO) GetColumns() []string {
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
