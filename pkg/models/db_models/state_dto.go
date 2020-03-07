package db_models

import "hydro_monitor/web_api/pkg/models/api_models"

type StateDTO struct {
	NodeId           string
	Name             string
	PhotosPerReading int
	ReadingInterval  int
	LowerLimit       float64
	UpperLimit       float64
	NextState        string
	PreviousState    string
}

func (s *StateDTO) GetColumns() []string {
	return nil
}

func (s *StateDTO) ConvertToAPIState() *api_models.State {
	return &api_models.State{
		Name:        s.Name,
		Interval:    s.ReadingInterval,
		UpperLimit:  s.UpperLimit,
		LowerLimit:  s.LowerLimit,
		PicturesNum: s.PhotosPerReading,
		Next:        s.NextState,
		Prev:        s.PreviousState,
	}
}
