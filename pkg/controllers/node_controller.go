package controllers

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
)

type NodeController interface {
	DeleteNode(c echo.Context) error
	PostNode(c echo.Context) error
	GetNodeByID(c echo.Context) error
	GetNodeManualReadingStatus(c echo.Context) error
	GetNodeConfiguration(c echo.Context) error
	UpdateNodeManualReading(c echo.Context) error
}

type nodeControllerImpl struct {
	nodeService services.NodeService
}

func (n *nodeControllerImpl) DeleteNode(c echo.Context) error {
	nodeId := c.Param("node_id")
	if err := n.nodeService.DeleteNode(nodeId); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (n *nodeControllerImpl) PostNode(c echo.Context) error {
	var node api_models.NodeDTO
	if err := c.Bind(&node); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err := n.nodeService.CreateNode(&node); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, node)
}

func (n *nodeControllerImpl) GetNodeManualReadingStatus(c echo.Context) error {
	nodeId := c.Param("node_id")
	respManualReading, err := n.nodeService.GetNodeManualReadingStatus(nodeId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, respManualReading)
}

func (n *nodeControllerImpl) UpdateNodeManualReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	var manualReading api_models.ManualReadingDTO
	if err := c.Bind(&manualReading); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	respManualReading, err := n.nodeService.UpdateNodeManualReading(nodeId, manualReading.ManualReading)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, respManualReading)
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
	nodeConfiguration, err := n.nodeService.GetNodeConfiguration(nodeId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nodeConfiguration)
}
