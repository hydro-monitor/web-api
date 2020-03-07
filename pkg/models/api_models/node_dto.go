package api_models

type NodeDTO struct {
	Id            string `json:"id,omitempty"`
	Description   string `json:"description"`
	ManualReading bool   `json:"manual_reading"`
}
