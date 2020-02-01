package api_models

type NodeDTO struct {
	Id            string `json:"id"`
	Description   string `json:"description"`
	Configuration string `json:"configuration"`
	State         string `json:"state"`
	ManualReading bool   `json:"manual_reading"`
}
