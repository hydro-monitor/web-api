package readings

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
	"time"
)

func GetAllNodeReadings(c echo.Context) error {
	nodeId := c.Param("id")
	readings, err := controllers.GetAllReadingsFromNode(nodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, readings)
}

func PostReading(c echo.Context) error {
	photo, _ := c.FormFile("photo")
	photo2, _ := photo.Open()
	raw_reading := c.FormValue("reading")
	reading := models.Reading{}
	_ = json.Unmarshal([]byte(raw_reading), &reading)
	reading.Photo = make([]byte, photo.Size)
	_, _ = photo2.Read(reading.Photo)
	if err := controllers.InsertReading(reading); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteReading(c echo.Context) error {
	nodeId := c.Param("id")
	timestamp, timeErr := time.Parse(time.RFC3339, c.Param("timestamp"))
	if timeErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := controllers.DeleteReading(timestamp, nodeId); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
