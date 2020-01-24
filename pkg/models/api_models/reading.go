package api_models

import "time"

type Reading struct {
	Time       time.Time `form:"timestamp"`
	WaterLevel float64   `form:"waterLevel"`
}
