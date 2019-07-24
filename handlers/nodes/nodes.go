package nodes

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/models"
	"net/http"
)

func GetNodes(c echo.Context) error {
	nodes := make([]models.Node, 2)
	nodes[0] = models.Node{"1", "Node 1"}
	nodes[1] = models.Node{"2", "Node 2"}
	return c.JSON(http.StatusOK, nodes)
}

func GetNode(c echo.Context) error {
	id := c.QueryParam("id")
	return c.JSON(http.StatusOK, models.Node{id, "A description"})
}
