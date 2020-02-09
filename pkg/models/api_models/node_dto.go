package api_models

type NodeDTO struct {
	Id            string `json:"id"`
	Description   string `json:"description"`
	ManualReading bool   `json:"manual_reading"`
}
