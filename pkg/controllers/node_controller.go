package controllers

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
)

type NodeController interface {
	GetNodeByID(c echo.Context) error
	GetNodeConfiguration(c echo.Context) error
}

type nodeControllerImpl struct {
	nodeService services.NodeService
}

func NewNodeController(nodeService services.NodeService) NodeController {
	return &nodeControllerImpl{nodeService}
}

func (n *nodeControllerImpl) GetNodeByID(c echo.Context) error {
	nodeId := c.Param("node_id")
	node, err := n.nodeService.GetNode(nodeId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, node)
}

func (n *nodeControllerImpl) GetNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	nodeConfiguration, _ := n.nodeService.GetNodeConfiguration(nodeId)
	return c.JSON(http.StatusOK, nodeConfiguration)
}
