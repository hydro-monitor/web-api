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
	UpdateNodeManualReading(c echo.Context) error
}

type nodeControllerImpl struct {
	nodeService services.NodeService
}

// CreateNodeConfiguration godoc
// @Summary Crea o actualiza la configuración para el nodo dado
// @Description Devuelve un mapa de estados (no un array como se ve a continuación) en donde la clave de cada uno es el nombre del mismo.
// @Tags nodes
// @Accept  json
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 200 {array} api_models.StateDTO
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/configuration [post]
func (n *nodeControllerImpl) CreateNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	configuration := make(map[string]*api_models.StateDTO)
	if err := c.Bind(&configuration); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := n.nodeService.CreateNodeConfiguration(nodeId, configuration); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, configuration)
}

// GetNodes godoc
// @Summary Obtiene todos los nodos
// @Tags nodes
// @Produce  json
// @Success 200 {array} api_models.NodeDTO
// @Failure 500 {object} echo.HTTPError
// @Router /nodes [get]
func (n *nodeControllerImpl) GetNodes(c echo.Context) error {
	nodes, err := n.nodeService.GetNodes()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, nodes)
}

// DeleteNode godoc
// @Summary Borra un nodo
// @Tags nodes
// @Param node_id path string true "ID del nodo"
// @Success 204
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id} [delete]
func (n *nodeControllerImpl) DeleteNode(c echo.Context) error {
	nodeId := c.Param("node_id")
	if err := n.nodeService.DeleteNode(nodeId); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// PostNode godoc
// @Summary Crea un nodo
// @Tags nodes
// @Accept  json
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 201 {object} api_models.NodeDTO
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes [post]
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

// GetNodeManualReadingStatus godoc
// @Summary Obtiene el estado de medición manual de un nodo
// @Tags nodes
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 200 {object} api_models.ManualReadingDTO
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/manual-reading [get]
func (n *nodeControllerImpl) GetNodeManualReadingStatus(c echo.Context) error {
	nodeId := c.Param("node_id")
	respManualReading, err := n.nodeService.GetNodeManualReadingStatus(nodeId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, respManualReading)
}

// GetNodeManualReadingStatus godoc
// @Summary Actualiza el estado de medición manual de un nodo
// @Tags nodes
// @Accept  json
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 200 {object} api_models.ManualReadingDTO
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/manual-reading [put]
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

// GetNodeByID godoc
// @Summary Obtiene la información completa de un nodo
// @Tags nodes
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 200 {object} api_models.NodeDTO
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id} [get]
func (n *nodeControllerImpl) GetNodeByID(c echo.Context) error {
	nodeId := c.Param("node_id")
	node, err := n.nodeService.GetNode(nodeId)
	if err != nil {
		return err.ToHTTPError()
	}
	return c.JSON(http.StatusOK, node)
}

// GetNodeConfiguration godoc
// @Summary Obtiene la configuración de un nodo
// @Description Devuelve un mapa de estados (no un array como se ve a continuación) en donde la clave de cada uno es el nombre del mismo.
// @Tags nodes
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Success 200 {array} api_models.StateDTO
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/configuration [get]
func (n *nodeControllerImpl) GetNodeConfiguration(c echo.Context) error {
	nodeId := c.Param("node_id")
	configuration, err := n.nodeService.GetNodeConfiguration(nodeId)
	if err != nil {
		return err.ToHTTPError()
	}
	return c.JSON(http.StatusOK, configuration)
}
