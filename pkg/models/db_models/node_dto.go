package db_models

import "hydro_monitor/web_api/pkg/models/api_models"

type NodeDTO struct {
	Id            string `db:"id"`
	Description   string `db:"description"`
	Configuration string `db:"configuration"`
	State         string `db:"state"`
	ManualReading bool   `db:"manual_reading"`
}

func (n *NodeDTO) GetColumns() []string {
	return nil
}

func (n *NodeDTO) ToAPINodeDTO() *api_models.NodeDTO {
	return &api_models.NodeDTO{
		Id:            n.Id,
		Description:   n.Description,
		Configuration: n.Configuration,
		State:         n.State,
		ManualReading: n.ManualReading,
	}
}
