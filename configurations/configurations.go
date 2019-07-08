package configurations

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Configuration struct {
	nodeId string
	name string
	photosPerReading int
	msBetweenReadings int
	waterLevelLimitForGoingToPreviousState int
	waterLevelLimitForGoingToNextState int
	previousState string
	nextState string
}

func GetNodeConfiguration(c echo.Context) error {
	nodeId := c.QueryParam("id")
	// TODO Select configuration from node
	msg := "Configuration for node " + nodeId
	return c.String(http.StatusOK, msg)
}

func PostNodeConfiguration(c echo.Context) error {
	nodeId := c.QueryParam("id")
	// TODO Update configuration
	msg := "Updated configuration for node " + nodeId
	return c.String(http.StatusOK, msg)
}
