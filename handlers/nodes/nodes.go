package nodes

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/db_driver/controllers"
	"hydro_monitor/web_api/models"
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
		return c.JSON(http.StatusInternalServerError, "Error when getting node with id:"+id+"with error: "+err.Error())
	}
	return c.JSON(http.StatusOK, node)
}

func PostNode(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	if err := controllers.InsertNode(models.Node{Id: m["id"].(string), Description: m["description"].(string)}); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteNode(c echo.Context) error {
	id := c.Param("id")
	if err := controllers.DeleteNode(id); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
