package nodes

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/db_driver/controllers"
	"hydro_monitor/web_api/pkg/models"
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
	id := c.Param("id")
	node, err := controllers.GetNodeByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error when getting node with id: "+id+" with error: "+err.Error())
	}
	return c.JSON(http.StatusOK, node)
}

func PostNode(c echo.Context) error {
	var node models.Node
	if err := c.Bind(&node); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := controllers.InsertNode(node); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteNode(c echo.Context) error {
	id := c.Param("id")
	if err := controllers.DeleteNode(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
