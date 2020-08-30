package controllers

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api_models"
	"net/http"
)

type HealthController interface {
	GetHealthStatus(c echo.Context) error
}

type healthControllerImpl struct {
}

func (h *healthControllerImpl) GetHealthStatus(c echo.Context) error {
	healthDTO := api_models.HealthDTO{Status: "pass"}
	return c.JSON(http.StatusOK, healthDTO)
}

func NewHealthController() HealthController {
	return &healthControllerImpl{}
}
