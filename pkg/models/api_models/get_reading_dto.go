package api_models

import "time"

type GetReadingDTO struct {
	ReadingId     string    `json:"readingId,omitempty" example:"00336270-8191-11ea-a43d-0242ac120003"`
	ReadingTime   time.Time `json:"readingTime" example:"2020-04-26T19:47:53.391Z"`
	WaterLevel    float64   `json:"waterLevel" example:"60"`
	ManualReading bool      `json:"manualReading" example:"false"`
}
