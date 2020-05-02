package api_models

type StateDTO struct {
	Interval    int     `json:"interval" example:"25"`
	UpperLimit  float64 `json:"upperLimit" example:"1"`
	LowerLimit  float64 `json:"lowerLimit" example:"0.5"`
	PicturesNum int     `json:"picturesNum" example:"1"`
	Next        string  `json:"next" example:"Alto"`
	Prev        string  `json:"prev" example:"Bajo"`
}
