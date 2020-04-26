package api_models

import "time"

type ReadingDTO struct {
	Time       time.Time `form:"timestamp" example:"2020-04-26T19:47:53.391Z"`
	WaterLevel float64   `form:"waterLevel" example:"60"`
}
