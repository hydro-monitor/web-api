package services

import (
	"hydro_monitor/web_api/pkg/controllers"
	"io/ioutil"
)

type ReadingsService interface {
	GetReading(id string)
	CreateReading(reading controllers.Reading) error
}

type readingsServiceImpl struct {
}

func (r readingsServiceImpl) CreateReading(reading controllers.Reading) error {
	return ioutil.WriteFile(reading.Time.String(), reading.Picture, 0644)
}

func (r readingsServiceImpl) GetReading(id string) {
	panic("implement me")
}

func NewReadingsService() ReadingsService {
	return &readingsServiceImpl{}
}
