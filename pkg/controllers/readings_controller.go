package controllers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/services"
	"io"
	"net/http"
	"strconv"
)

type ReadingsController interface {
	CreateReading(c echo.Context) error
	AddPhotoToReading(c echo.Context) error
	GetNodeReadings(c echo.Context) error
	GetNodeReading(c echo.Context) error
	GetNodesLastReading(c echo.Context) error
	GetReadingPhoto(c echo.Context) error
	DeleteReading(c echo.Context) error
}

type readingsControllerImpl struct {
	nodesService    services.NodeService
	readingsService services.ReadingsService
}

// DeleteReading godoc
// @Summary Borra una medición
// @Tags readings
// @Param node_id path string true "ID del nodo"
// @Param reading_id path string true "ID de la medición"
// @Success 204 "No Content"
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings/{reading_id} [delete]
func (r *readingsControllerImpl) DeleteReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	readingId := c.Param("reading_id")
	if nodeId == "" || readingId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Node ID and reading ID can't be null")
	}
	if err := r.readingsService.DeleteReading(nodeId, readingId); err != nil {
		return err.ToHTTPError()
	}
	return c.NoContent(http.StatusNoContent)
}

// GetNodesLastReading godoc
// @Summary Obtiene la última medición de todos los nodos
// @Tags readings
// @Param node_id path string true "ID del nodo"
// @Success 200 {object} map[string]api_models.GetReadingDTO
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/last-reading [get]
func (r *readingsControllerImpl) GetNodesLastReading(c echo.Context) error {
	nodes, err := r.nodesService.GetNodes()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	lastReadings, err2 := r.readingsService.GetNodesLastReading(nodes)
	if err2 != nil {
		return err2
	}
	return c.JSON(http.StatusOK, lastReadings)
}

// AddPhotoToReading godoc
// @Summary Agrega una foto a la medición
// @Tags readings
// @Accept jpeg
// @Param node_id path string true "ID del nodo"
// @Param reading_id path string true "ID de la medición"
// @Param picture formData string true "Foto de la medición"
// @Success 201 "Created"
// @Failure 400 {object} echo.HTTPError
// @Failure 422 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings/{reading_id}/photos [post]
func (r *readingsControllerImpl) AddPhotoToReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	readingId := c.Param("reading_id")
	if nodeId == "" || readingId == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Node ID and reading ID can't be null")
	}
	photo := new(api_models.PhotoDTO)
	photo.ReadingId = readingId
	if err := c.Bind(photo); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	photoFile, err := r.extractPicture(c)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}
	photo.Photo = photoFile
	if err := r.readingsService.AddPhotoToReading(nodeId, photo); err != nil {
		return err.ToHTTPError()
	}
	return c.NoContent(http.StatusCreated)
}

// GetNodeReading godoc
// @Summary Obtiene los datos de una medición
// @Tags readings
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Param reading_id path string true "ID de la medición"
// @Success 200 {object} api_models.GetReadingDTO
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings/{reading_id} [get]
func (r *readingsControllerImpl) GetNodeReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	readingId := c.Param("reading_id")
	apiReading, err := r.readingsService.GetNodeReading(nodeId, readingId)
	if err != nil {
		return err.ToHTTPError()
	}
	return c.JSON(http.StatusOK, apiReading)
}

// GetNodeReadings godoc
// @Summary Obtiene las mediciones de un nodo
// @Tags readings
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Param page_size query int false "Cantidad máxima de mediciones por página"
// @Param page_state query string false "String en base 64 que contiene el estado de pagina. Utilizado para traer la próxima página"
// @Success 200 {array} api_models.GetReadingDTO
// @Header 200 {string} Page-State "Page state"
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings [get]
func (r *readingsControllerImpl) GetNodeReadings(c echo.Context) error {
	nodeId := c.Param("node_id")
	rawPageSize := c.QueryParam("page_size")
	rawPageState := c.QueryParam("page_state")
	var pageSize int
	var pageState []byte
	if rawPageSize != "" {
		tempPageSize, err := strconv.Atoi(c.QueryParam("page_size"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Page size must be a number")
		} else {
			pageSize = tempPageSize
		}
	}
	if rawPageState != "" {
		tempPageState, err := base64.StdEncoding.DecodeString(rawPageState)
		if err != nil {
			log.Errorf(fmt.Sprintf("Bad page state: %s, error: %s", rawPageState, err.Error()))
			return echo.NewHTTPError(http.StatusBadRequest, "Ill formed page state")
		} else {
			pageState = tempPageState
		}
	}
	paginatedDTO, err := r.readingsService.GetNodeReadings(nodeId, pageState, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	c.Response().Header().Set("Page-State", base64.StdEncoding.EncodeToString(paginatedDTO.PageState))
	return c.JSON(http.StatusOK, paginatedDTO.Elements)
}

// GetReadingPhoto godoc
// @Summary Obtiene la foto de una medición
// @Tags readings
// @Produce  jpeg
// @Param node_id path string true "ID del nodo"
// @Param reading_id path string true "ID de la medición"
// @Success 200 "Ok"
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings/{reading_id}/photos [get]
func (r *readingsControllerImpl) GetReadingPhoto(c echo.Context) error {
	readingId := c.Param("reading_id")
	photo, err := r.readingsService.GetReadingPhoto(readingId, 0)
	if err != nil {
		return err.ToHTTPError()
	}
	return c.Blob(http.StatusOK, "image/jpeg", photo.Picture)
}

// CreateReading godoc
// @Summary Crea una medición
// @Tags readings
// @Accept json
// @Produce  json
// @Param node_id path string true "ID del nodo"
// @Param reading body api_models.ReadingDTO true "Datos de la medición"
// @Success 200 {object} api_models.GetReadingDTO
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /nodes/{node_id}/readings [post]
func (r *readingsControllerImpl) CreateReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	reading := new(api_models.ReadingDTO)
	if err := c.Bind(reading); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	newReading, err := r.readingsService.CreateReading(nodeId, reading)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newReading)
}

func (r *readingsControllerImpl) extractPicture(c echo.Context) ([]byte, error) {
	file, err := c.FormFile("picture")
	if err != nil {
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NewReadingsController(nodesService services.NodeService, readingsService services.ReadingsService) ReadingsController {
	return &readingsControllerImpl{nodesService: nodesService, readingsService: readingsService}
}
