package db_models

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/models/api_models"
	"time"
)

type Reading struct {
	NodeId        string
	ReadingId     gocql.UUID
	ReadingTime   time.Time
	WaterLevel    float64
	ManualReading bool
}

func (r *Reading) GetColumns() []string {
	return nil
}

func (r *Reading) ConvertToAPIGetReading() *api_models.GetReadingDTO {
	return &api_models.GetReadingDTO{
		ReadingId:     r.ReadingId.String(),
		ReadingTime:   r.ReadingTime,
		WaterLevel:    r.WaterLevel,
		ManualReading: r.ManualReading,
	}
}
