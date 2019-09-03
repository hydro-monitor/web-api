package manual_readings

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
)

func GetManualReadingFromNode(c echo.Context) error {
	nodeId := c.Param("id")
	reading, err := controllers.GetManualReading(nodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			"Error when getting manual reading with id: "+nodeId+" with error: "+err.Error())
	}
	return c.JSON(http.StatusOK, reading)
}

func UpdateManualReading(c echo.Context) error {
	var reading models.ManualReading
	if err := c.Bind(&reading); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if applied, err := controllers.UpdateManualReading(reading); err != nil || !applied {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
