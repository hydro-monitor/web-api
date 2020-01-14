package controllers

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api"
	"hydro_monitor/web_api/pkg/services"
	"io/ioutil"
	"net/http"
)

type NodeController interface {
	GetNodeByID(c echo.Context) error
	GetNodeConfiguration(c echo.Context) error
	CreateReading(c echo.Context) error
}

type nodeControllerImpl struct {
	nodeService services.NodeService
}

func (n *nodeControllerImpl) CreateReading(c echo.Context) error {
	reading := new(api.Reading)
	if err := c.Bind(reading); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err := ioutil.WriteFile(reading.Time.String(), reading.Picture, 0644)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
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
