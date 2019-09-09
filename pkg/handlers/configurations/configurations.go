package configurations

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
)

func GetNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("id")
	configuration, err := controllers.GetNodeConfiguration(nodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, configuration)
}

func PutNodeConfiguration(c echo.Context) error {
	var configuration models.Configuration
	if err := c.Bind(&configuration); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controllers.UpdateNodeConfiguration(configuration); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
