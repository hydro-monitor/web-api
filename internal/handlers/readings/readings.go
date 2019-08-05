package readings

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
	"time"
)

func GetAllNodeReadings(c echo.Context) error {
	nodeId := c.QueryParam("id")
	readings, err := controllers.GetAllReadingsFromNode(nodeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, readings)
}

func PostReading(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	timestamp, err := time.Parse(time.RFC3339, m["timestamp"].(string))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if applied, err := controllers.InsertReading(
		models.Reading{
			Timestamp:  timestamp,
			NodeId:     m["id"].(string),
			WaterLevel: m["waterlevel"].(float64),
			Photo:      m["photo"].([]byte)}); err != nil || applied == false {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteReading(c echo.Context) error {
	nodeId := c.QueryParam("id")
	timestamp, timeErr := time.Parse(time.RFC3339, c.QueryParam("timestamp"))
	if timeErr != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if applied, err := controllers.DeleteReading(timestamp, nodeId); err != nil || applied == false {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
