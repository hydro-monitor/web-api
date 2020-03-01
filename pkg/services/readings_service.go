package services

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
)

type ReadingsService interface {
	CreateReading(nodeId string, reading *api_models.Reading) (*api_models.GetReadingDTO, error)
	AddPhotoToReading(photoDTO *api_models.PhotoDTO) (*api_models.PhotoMetadataDTO, error)
	GetNodeReadings(nodeId string) ([]*api_models.GetReadingDTO, error)
	GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, ServiceError)
	GetReadingPhoto(readingId string, number int) (*db_models.Photo, ServiceError)
}

type readingsServiceImpl struct {
	readingsRepository repositories.Repository
	photosRepository   repositories.Repository
}

func (r *readingsServiceImpl) AddPhotoToReading(photoDTO *api_models.PhotoDTO) (*api_models.PhotoMetadataDTO, error) {
	readingUUID, err := gocql.ParseUUID(photoDTO.ReadingId)
	if err != nil {
		return nil, err
	}
	dbPhoto := &db_models.Photo{
		ReadingTime: readingUUID,
		Number:      photoDTO.PhotoNumber,
		Picture:     photoDTO.Photo,
	}
	if err := r.photosRepository.Insert(dbPhoto); err != nil {
		return nil, err
	}
	apiPhotoMetadata := api_models.PhotoMetadataDTO{
		ReadingId:   photoDTO.ReadingId,
		PhotoNumber: photoDTO.PhotoNumber,
	}
	return &apiPhotoMetadata, nil
}

func (r *readingsServiceImpl) GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, ServiceError) {
	readingUUID, err := gocql.ParseUUID(readingId)
	if err != nil {
		return nil, NewBadReadingTimeError("Incorrect reading time (bad format)", err)
	}
	dbReading := db_models.Reading{
		NodeId:    nodeId,
		ReadingId: readingUUID,
	}
	err = r.readingsRepository.Get(&dbReading)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node reading not found", err)
		}
		return nil, NewGenericServiceError("Server error when getting node reading", err)
	}
	apiReading := api_models.GetReadingDTO{WaterLevel: dbReading.WaterLevel}
	return &apiReading, nil
}

func (r *readingsServiceImpl) GetReadingPhoto(readingId string, number int) (*db_models.Photo, ServiceError) {
	readingUUID, err := gocql.ParseUUID(readingId)
	if err != nil {
		return nil, NewBadReadingTimeError("Incorrect reading time (bad format)", err)
	}
	dbPhoto := db_models.Photo{ReadingTime: readingUUID}
	err = r.photosRepository.Get(&dbPhoto)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Reading photo not found", err)
		}
		return nil, NewGenericServiceError("Server error when getting reading photo", err)
	}
	return &dbPhoto, nil
}

func (r *readingsServiceImpl) GetNodeReadings(nodeId string) ([]*api_models.GetReadingDTO, error) {
	readings := db_models.NewReadingsDTO(nodeId)
	if err := r.readingsRepository.Select(readings); err != nil {
		return nil, err
	}
	return readings.ConvertToAPIGetReadings(), nil
}

func (r *readingsServiceImpl) CreateReading(nodeId string, reading *api_models.Reading) (*api_models.GetReadingDTO, error) {
	readingTimeUUID := gocql.UUIDFromTime(reading.Time)
	dbReading := &db_models.Reading{
		NodeId:      nodeId,
		ReadingId:   readingTimeUUID,
		ReadingTime: reading.Time,
		WaterLevel:  reading.WaterLevel,
	}
	if err := r.readingsRepository.Insert(dbReading); err != nil {
		return nil, err
	}
	return dbReading.ConvertToAPIGetReading(), nil
}

func NewReadingsService(client db.Client) ReadingsService {
	readingsRepository := repositories.NewReadingsRepository(client)
	photosRepository := repositories.NewPhotosRepository(client)
	return &readingsServiceImpl{readingsRepository: readingsRepository, photosRepository: photosRepository}
}
