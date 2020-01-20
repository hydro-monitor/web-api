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
	GetNodeReadings(nodeId string) error
	GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, error)
	GetReadingPhoto(readingId string, number int) (*db_models.Photo, error)
}

type readingsServiceImpl struct {
	readingsRepository repositories.Repository
	photosRepository   repositories.Repository
}

func (r *readingsServiceImpl) GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, error) {
	readingUUID, err := gocql.ParseUUID(readingId)
	if err != nil {
		return nil, err
	}
	dbReading := db_models.Reading{
		NodeId:      nodeId,
		ReadingTime: readingUUID,
	}
	err = r.readingsRepository.Get(&dbReading)
	if err != nil {
		return nil, err
	}
	apiReading := api_models.GetReadingDTO{
		WaterLevel: dbReading.WaterLevel,
		Pictures:   make([]int, 0),
	}
	return &apiReading, nil
	/*	dbPhoto := db_models.Photo{ReadingTime: readingUUID}
		err = r.photosRepository.Get(dbPhoto)*/
}

func (r *readingsServiceImpl) GetReadingPhoto(readingId string, number int) (*db_models.Photo, error) {
	readingUUID, err := gocql.ParseUUID(readingId)
	if err != nil {
		return nil, err
	}
	dbPhoto := db_models.Photo{ReadingTime: readingUUID}
	err = r.photosRepository.Get(&dbPhoto)
	return &dbPhoto, err
}

func (r *readingsServiceImpl) GetNodeReadings(nodeId string) error {
	reading := db_models.Reading{NodeId: nodeId}
	return r.readingsRepository.Get(&reading)
}

func (r *readingsServiceImpl) CreateReading(nodeId string, reading *api_models.Reading) (*api_models.GetReadingDTO, error) {
	readingTime := gocql.UUIDFromTime(reading.Time)
	dbReading := db_models.Reading{
		NodeId:      nodeId,
		ReadingTime: readingTime,
		WaterLevel:  reading.WaterLevel,
	}
	dbPhoto := db_models.Photo{
		ReadingTime: readingTime,
		Number:      0,
		Picture:     reading.Picture,
	}
	if err := r.readingsRepository.Insert(dbReading); err != nil {
		return nil, err
	}
	if err := r.photosRepository.Insert(dbPhoto); err != nil {
		return nil, err
	}
	apiReading := api_models.GetReadingDTO{
		NodeId:     dbReading.NodeId,
		ReadingId:  dbReading.ReadingTime.String(),
		WaterLevel: dbReading.WaterLevel,
		Pictures:   make([]int, 0),
	}
	apiReading.Pictures = append(apiReading.Pictures, dbPhoto.Number)
	return &apiReading, nil
}

func NewReadingsService(client db.Client) ReadingsService {
	readingsRepository := repositories.NewReadingsRepository(client)
	photosRepository := repositories.NewPhotosRepository(client)
	return &readingsServiceImpl{readingsRepository: readingsRepository, photosRepository: photosRepository}
}
