package services

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
)

type ReadingsService interface {
	CreateReading(nodeId string, reading *api_models.ReadingDTO) (*api_models.GetReadingDTO, error)
	AddPhotoToReading(photoDTO *api_models.PhotoDTO) (*api_models.PhotoMetadataDTO, error)
	GetNodeReadings(nodeId string, pageState []byte, pageSize int) ([]*api_models.GetReadingDTO, ServiceError)
	GetNodeReading(nodeId string, readingId string) (*api_models.GetReadingDTO, ServiceError)
	GetNodesLastReading(nodes []*api_models.NodeDTO) (map[string]*api_models.GetReadingDTO, ServiceError)
	GetReadingPhoto(readingId string, number int) (*db_models.Photo, ServiceError)
}

type readingsServiceImpl struct {
	nodesRepository    repositories.Repository
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
	return dbReading.ConvertToSingleAPIGetReading(), nil
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
			return nil, NewNotFoundError("ReadingDTO photo not found", err)
		}
		return nil, NewGenericServiceError("Server error when getting reading photo", err)
	}
	return &dbPhoto, nil
}

func (r *readingsServiceImpl) GetNodeReadings(nodeId string, pageState []byte, pageSize int) ([]*api_models.GetReadingDTO, ServiceError) {
	readings := db_models.NewReadingsDTO(nodeId)
	if err := r.readingsRepository.Select(readings, pageState, pageSize); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node readings not found", err)
		}
		return nil, NewGenericServiceError("Server error when getting node readings", err)
	}
	return readings.ConvertToAPIGetReadings(), nil
}

func (r *readingsServiceImpl) GetNodesLastReading(nodes []*api_models.NodeDTO) (map[string]*api_models.GetReadingDTO, ServiceError) {
	lastReadings := make(map[string]*api_models.GetReadingDTO)
	for _, node := range nodes {
		readings, err := r.GetNodeReadings(node.Id, nil, 1)
		if err != nil {
			return nil, err
		}
		if len(readings) == 1 {
			lastReadings[node.Id] = readings[0]
		} else {
			// TODO ver si hacer esto o directamente no devolver el nodo
			lastReadings[node.Id] = nil
		}
	}
	return lastReadings, nil
}

func (r *readingsServiceImpl) CreateReading(nodeId string, reading *api_models.ReadingDTO) (*api_models.GetReadingDTO, error) {
	if err := r.nodesRepository.Get(&db_models.NodeDTO{Id: nodeId}); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("Node not found", err)
		}
		return nil, NewGenericServiceError("Server error when trying to create new reading", err)
	}
	readingTimeUUID := gocql.UUIDFromTime(reading.Time)
	dbReading := &db_models.Reading{
		NodeId:        nodeId,
		ReadingId:     readingTimeUUID,
		ReadingTime:   reading.Time,
		WaterLevel:    reading.WaterLevel,
		ManualReading: reading.ManualReading,
	}
	if err := r.readingsRepository.Insert(dbReading); err != nil {
		return nil, err
	}
	return dbReading.ConvertToAPIGetReading(), nil
}

func NewReadingsService(nodesRepository repositories.Repository, photosRepository repositories.Repository, readingsRepository repositories.Repository) ReadingsService {
	return &readingsServiceImpl{
		nodesRepository:    nodesRepository,
		readingsRepository: readingsRepository,
		photosRepository:   photosRepository,
	}
}
