package nodes

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/db_driver/controllers"
	"net/http"
)

func GetNodes(c echo.Context) error {
	nodes, err := controllers.GetAllNodes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nodes)
}

func GetNode(c echo.Context) error {
	id := c.QueryParam("id")
	node, err := controllers.GetNodeByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, node)
}
