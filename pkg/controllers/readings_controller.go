package controllers

import (
	"github.com/labstack/echo"
	"hydro_monitor/web_api/pkg/models/api"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
)

type ReadingsController interface {
	CreateReading(c echo.Context) error
}

type readingsControllerImpl struct {
	service services.ReadingsService
}

func (r *readingsControllerImpl) GetNodeByID(c echo.Context) error {
	panic("implement me")
}

func (r *readingsControllerImpl) CreateReading(c echo.Context) error {
	reading := new(api.Reading)
	if err := c.Bind(reading); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusCreated)
}

func NewReadingsController(service services.ReadingsService) ReadingsController {
	return &readingsControllerImpl{service: service}
}
