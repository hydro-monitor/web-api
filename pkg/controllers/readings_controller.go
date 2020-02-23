package controllers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/services"
	"io"
	"net/http"
)

type ReadingsController interface {
	CreateReading(c echo.Context) error
	AddPhotoToReading(c echo.Context) error
	GetNodeReadings(c echo.Context) error
	GetNodeReading(c echo.Context) error
	GetReadingPhoto(c echo.Context) error
}

type readingsControllerImpl struct {
	service services.ReadingsService
}

func (r *readingsControllerImpl) AddPhotoToReading(c echo.Context) error {
	readingId := c.Param("reading_id")
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
	r.service.AddPhotoToReading(photo)
	return c.NoContent(http.StatusCreated)
}

func (r *readingsControllerImpl) GetNodeReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	readingId := c.Param("reading_id")
	apiReading, err := r.service.GetNodeReading(nodeId, readingId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, apiReading)
}

func (r *readingsControllerImpl) GetNodeReadings(c echo.Context) error {
	nodeId := c.Param("node_id")
	getReadings, err := r.service.GetNodeReadings(nodeId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, getReadings)
}

func (r *readingsControllerImpl) GetReadingPhoto(c echo.Context) error {
	readingId := c.Param("reading_id")
	photo, err := r.service.GetReadingPhoto(readingId, 0)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.Blob(http.StatusOK, "image/jpeg", photo.Picture)
}

func (r *readingsControllerImpl) CreateReading(c echo.Context) error {
	nodeId := c.Param("node_id")
	reading := new(api_models.Reading)
	if err := c.Bind(reading); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	newReading, err := r.service.CreateReading(nodeId, reading)
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

func NewReadingsController(service services.ReadingsService) ReadingsController {
	return &readingsControllerImpl{service: service}
}
