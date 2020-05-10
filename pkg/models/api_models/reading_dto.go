package api_models

import "time"

type ReadingDTO struct {
	Time          time.Time `json:"timestamp" form:"timestamp" example:"2020-04-26T19:47:53.391Z"`
	WaterLevel    float64   `json:"waterLevel" form:"waterLevel" example:"60"`
	ManualReading bool      `json:"manualReading" form:"manualReading" example:"false"`
}
