package services

import (
	"hydro_monitor/web_api/pkg/models/api"
	"io/ioutil"
)

type ReadingsService interface {
	CreateReading(reading api.Reading) error
}

type readingsServiceImpl struct {
}

func (r readingsServiceImpl) CreateReading(reading api.Reading) error {
	return ioutil.WriteFile(reading.Time.String(), reading.Picture, 0644)
}

func NewReadingsService() ReadingsService {
	return &readingsServiceImpl{}
}
