package api_models

type GetReadingDTO struct {
	NodeId     string  `json:"nodeId"`
	ReadingId  string  `json:"readingId"`
	WaterLevel float64 `json:"waterLevel"`
	Pictures   []int   `json:"pictures"`
}
