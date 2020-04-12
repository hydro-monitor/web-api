package controllers

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
)

type NodeController interface {
	CreateNodeConfiguration(c echo.Context) error
	DeleteNode(c echo.Context) error
	PostNode(c echo.Context) error
	GetNodes(c echo.Context) error
	GetNodeByID(c echo.Context) error
	GetNodeManualReadingStatus(c echo.Context) error
	GetNodeConfiguration(c echo.Context) error
	UpdateNodeConfiguration(c echo.Context) error
	UpdateNodeManualReading(c echo.Context) error
}

type nodeControllerImpl struct {
	nodeService services.NodeService
}

func (n *nodeControllerImpl) UpdateNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	var states []*api_models.State
	if err := c.Bind(&states); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := n.nodeService.CreateNodeConfiguration(nodeId, states); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, states)
}

func (n *nodeControllerImpl) CreateNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	var states []*api_models.State
	if err := c.Bind(&states); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// TODO decide whether having two endpoints for creating/updating node configurations is necessary
	if err := n.nodeService.CreateNodeConfiguration(nodeId, states); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, states)
}

func (n *nodeControllerImpl) GetNodes(c echo.Context) error {
	nodes, err := n.nodeService.GetNodes()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nodes)
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
		return err.ToHTTPError()
	}
	return c.JSON(http.StatusOK, node)
}

func (n *nodeControllerImpl) GetNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	_, err := n.nodeService.GetNodeConfiguration(nodeId)
	if err != nil {
		return err.ToHTTPError()
	}
	state1 := api_models.StateDTO{
		Interval:    25,
		UpperLimit:  1,
		LowerLimit:  -999,
		PicturesNum: 1,
		Next:        "2",
		Prev:        "",
	}
	state2 := api_models.StateDTO{
		Interval:    10,
		UpperLimit:  2,
		LowerLimit:  1,
		PicturesNum: 1,
		Next:        "3",
		Prev:        "1",
	}
	state3 := api_models.StateDTO{
		Interval:    5,
		UpperLimit:  999,
		LowerLimit:  2,
		PicturesNum: 1,
		Next:        "",
		Prev:        "2",
	}
	stateMap := make(map[string]*api_models.StateDTO)
	stateMap["1"] = &state1
	stateMap["2"] = &state2
	stateMap["3"] = &state3
	return c.JSON(http.StatusOK, stateMap)
}
