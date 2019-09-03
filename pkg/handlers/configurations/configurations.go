package configurations

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"net/http"
)

func GetNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("id")
	configuration, err := controllers.GetNodeConfiguration(nodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, configuration)
}

func PutNodeConfiguration(c echo.Context) error {
	nodeId := c.QueryParam("id")
	// TODO Update configuration
	msg := "Updated configuration for node " + nodeId
	return c.String(http.StatusOK, msg)
}
