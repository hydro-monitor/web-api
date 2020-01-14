package controllers

import (
	"github.com/labstack/echo"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
	"time"
)

type Reading struct {
	Time       time.Time `form:"timestamp"`
	WaterLevel float64   `form:"waterLevel"`
	Picture    []byte    `form:"picture"`
}

type ReadingsController interface {
	GetReading(c echo.Context) error
	GetReadingsFromNode(c echo.Context) error
	CreateReading(c echo.Context) error
}

type readingsControllerImpl struct {
	service services.ReadingsService
}

func (r readingsControllerImpl) GetReading(c echo.Context) error {
	panic("implement me")
}

func (r readingsControllerImpl) GetReadingsFromNode(c echo.Context) error {
	panic("implement me")
}

func (r readingsControllerImpl) CreateReading(c echo.Context) error {
	reading := new(Reading)
	if err := c.Bind(reading); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}

func NewReadingsController(service services.ReadingsService) ReadingsController {
	return readingsControllerImpl{service: service}
}
