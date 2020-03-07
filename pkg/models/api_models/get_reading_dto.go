package api_models

import "time"

type GetReadingDTO struct {
	ReadingId     string    `json:"readingId"`
	ReadingTime   time.Time `json:"readingTime"`
	WaterLevel    float64   `json:"waterLevel"`
	ManualReading bool      `json:"manualReading"`
}
