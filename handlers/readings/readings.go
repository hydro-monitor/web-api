package readings

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Reading struct {
	timestamp time.Time
	nodeId    string
	waterLevel float64
	photo []byte
}

func GetAllNodeReadings(c echo.Context) error {
	nodeId := c.QueryParam("id")
	// TODO Select action to Cassandra Driver
	msg := "Queried readings from node " + nodeId
	return c.String(http.StatusOK, msg)
}

func DeleteReading(c echo.Context) error {
	nodeId := c.QueryParam("id")
	timestamp := c.QueryParam("timestamp")
	// TODO Delete action to Cassandra Driver
	msg := "Deleted reading " + timestamp + " from node " + nodeId
	return c.String(http.StatusOK, msg)
}
