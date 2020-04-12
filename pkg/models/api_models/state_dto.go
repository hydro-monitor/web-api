package api_models

type StateDTO struct {
	Interval    int     `json:"interval"`
	UpperLimit  float64 `json:"upperLimit"`
	LowerLimit  float64 `json:"lowerLimit"`
	PicturesNum int     `json:"picturesNum"`
	Next        string  `json:"next"` // State name (key)
	Prev        string  `json:"prev"` // State name (key)
}
