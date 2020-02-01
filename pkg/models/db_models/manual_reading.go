package db_models

import "hydro_monitor/web_api/pkg/models/api_models"

var columns = []string{"manual_reading"}

type ManualReadingDTO struct {
	NodeId        string `db:"id"`
	ManualReading bool   `db:"manual_reading"`
}

func (m *ManualReadingDTO) GetColumns() []string {
	return columns
}

func (m *ManualReadingDTO) ToAPIManualReadingDTO() *api_models.ManualReadingDTO {
	return &api_models.ManualReadingDTO{ManualReading: m.ManualReading}
}
