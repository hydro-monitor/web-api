package db_models

import (
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/models/api_models"
)

type ReadingsDTO struct {
	nodeId   string
	Readings []*Reading
}

func (r *ReadingsDTO) GetColumns() []string {
	return nil
}

func (r *ReadingsDTO) GetBindMap() qb.M {
	return qb.M{"node_id": r.nodeId}
}

func (r *ReadingsDTO) GetArgs() interface{} {
	return &r.Readings
}

func (r *ReadingsDTO) ConvertToAPIGetReadings() []*api_models.GetReadingDTO {
	var getReadings []*api_models.GetReadingDTO
	for _, r := range r.Readings {
		getReadings = append(getReadings, r.ConvertToAPIGetReading())
	}
	return getReadings
}

func NewReadingsDTO(nodeId string) *ReadingsDTO {
	readings := make([]*Reading, 0)
	return &ReadingsDTO{
		nodeId:   nodeId,
		Readings: readings,
	}
}
