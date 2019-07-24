package nodes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Node struct {
	Id          string `json:"id" form:"id" query:"id"`
	Description string `json:"description" form:"description" query:"description"`
}

func GetNodes(c echo.Context) error {
	nodes := make([]Node, 2)
	nodes[0] = Node{"1", "Node 1"}
	nodes[1] = Node{"2", "Node 2"}
	return c.JSON(http.StatusOK, nodes)
}

func GetNode(c echo.Context) error {
	id := c.QueryParam("id")
	return c.JSON(http.StatusOK, Node{id, "A description"})
}
