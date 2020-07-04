package services_test

import (
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/mocks"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/services"
	"testing"
	"time"
)

type ReadingsServiceTestSuite struct {
	suite.Suite
	readingsService        services.ReadingsService
	nodesRepositoryMock    *mocks.Repository
	photosRepositoryMock   *mocks.Repository
	readingsRepositoryMock *mocks.Repository
}

func (suite *ReadingsServiceTestSuite) SetupTest() {
	suite.nodesRepositoryMock = new(mocks.Repository)
	suite.photosRepositoryMock = new(mocks.Repository)
	suite.readingsRepositoryMock = new(mocks.Repository)
	suite.readingsService = services.NewReadingsService(suite.nodesRepositoryMock, suite.photosRepositoryMock, suite.readingsRepositoryMock)
}

func (suite *ReadingsServiceTestSuite) TestGetNodeReading() {
	nodeId := "lujan-1"
	readingId := "00336270-8191-11ea-a43d-0242ac120003"
	// readingUUID, _ := gocql.ParseUUID(readingId)

	suite.readingsRepositoryMock.On("Get", mock.Anything).Return(nil)

	_, err := suite.readingsService.GetNodeReading(nodeId, readingId)

	suite.Nil(err)
	// suite.EqualValues(readingUUID.Time(), res.ReadingTime)
}

func (suite *ReadingsServiceTestSuite) TestCreateReading() {
	nodeId := "lujan-1"
	reading := &api_models.ReadingDTO{Time: time.Now(), WaterLevel: 55.5}
	expectedReading := &api_models.GetReadingDTO{
		ReadingTime:   reading.Time,
		WaterLevel:    reading.WaterLevel,
		ManualReading: false,
	}

	suite.nodesRepositoryMock.On("Get", &db_models.NodeDTO{Id: &nodeId}).Return(nil)
	suite.readingsRepositoryMock.On("Insert", mock.Anything).Return(nil)

	res, err := suite.readingsService.CreateReading(nodeId, reading)

	suite.Nil(err)
	suite.Equal(expectedReading.ReadingTime, res.ReadingTime)
	suite.Equal(expectedReading.WaterLevel, res.WaterLevel)
	suite.Equal(expectedReading.ManualReading, res.ManualReading)
}

func (suite *ReadingsServiceTestSuite) TestCreateReadingWithNonExistingNode() {
	nodeId := "non-existing-node"
	reading := &api_models.ReadingDTO{Time: time.Now(), WaterLevel: 55.5}
	gocqlNotFoundError := gocql.ErrNotFound
	expectedError := services.NewNotFoundError("Node not found", gocqlNotFoundError)

	suite.nodesRepositoryMock.On("Get", &db_models.NodeDTO{Id: &nodeId}).Return(gocqlNotFoundError)

	res, err := suite.readingsService.CreateReading(nodeId, reading)

	suite.Equal((*api_models.GetReadingDTO)(nil), res)
	suite.Equal(expectedError, err)
}

func TestReadingsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ReadingsServiceTestSuite))
}
