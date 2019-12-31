package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type StateController interface {
	GetConfiguration(c echo.Context) error
}

type stateControllerImpl struct {
}

func NewStateController() StateController {
	return &stateControllerImpl{}
}

func (s *stateControllerImpl) GetConfiguration(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
