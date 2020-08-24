package services

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"

	"testing"
	"time"
)

type ReadingsServiceTestSuite struct {
	suite.Suite
	readingsService        ReadingsService
	nodesRepositoryMock    *repositories.RepositoryMock
	photosRepositoryMock   *repositories.RepositoryMock
	readingsRepositoryMock *repositories.RepositoryMock
}

func (suite *ReadingsServiceTestSuite) SetupTest() {
	suite.nodesRepositoryMock = new(repositories.RepositoryMock)
	suite.photosRepositoryMock = new(repositories.RepositoryMock)
	suite.readingsRepositoryMock = new(repositories.RepositoryMock)
	suite.readingsService = NewReadingsService(suite.nodesRepositoryMock, suite.photosRepositoryMock, suite.readingsRepositoryMock)
}

func (suite *ReadingsServiceTestSuite) TestGetNodeReading() {
	// nodeId := "lujan-1"
	// readingId := "00336270-8191-11ea-a43d-0242ac120003"
	// readingUUID, _ := gocql.ParseUUID(readingId)

	suite.readingsRepositoryMock.On("Get", mock.Anything).Return(nil)

	// _, err := suite.readingsService.GetNodeReading(nodeId, readingId)

	//suite.Nil(err)
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
	expectedError := NewNotFoundError("Node not found", gocqlNotFoundError)

	suite.nodesRepositoryMock.On("Get", &db_models.NodeDTO{Id: &nodeId}).Return(gocqlNotFoundError)

	res, err := suite.readingsService.CreateReading(nodeId, reading)

	suite.Equal((*api_models.GetReadingDTO)(nil), res)
	suite.Equal(expectedError, err)
}

func (suite *ReadingsServiceTestSuite) TestDeleteReading() {
	tests := []struct {
		name       string
		args       []string
		mockReturn []error
		want       ServiceError
	}{
		{
			name:       "when everything it's OK no error should be returned",
			args:       []string{"ok-node", "00336270-8191-11ea-a43d-0242ac120003"},
			mockReturn: []error{nil, nil},
			want:       nil,
		},
		{
			name:       "providing a non valid reading ID should return an invalid UUID error",
			args:       []string{"ok-node", "non-valid-uuid"},
			mockReturn: []error{nil, nil},
			want: NewGenericServiceError(
				"Error when trying to decode reading UUID",
				errors.New("invalid UUID \"non-valid-uuid\"")),
		},
		{
			name:       "receiving an error when attempting to delete the reading's photos should return an error",
			args:       []string{"reading-photo-repo-error-node", "00336270-8191-11ea-a43d-0242ac120003"},
			mockReturn: []error{errors.New("db error"), nil},
			want: NewGenericServiceError(
				"Error when trying to delete reading's photos",
				errors.New("db error")),
		},
		{
			name:       "receiving an error when attempting to delete the reading should return an error",
			args:       []string{"reading-repo-error-node", "00336270-8191-11ea-a43d-0242ac120003"},
			mockReturn: []error{nil, errors.New("db error")},
			want:       NewGenericServiceError("Error when trying to delete reading", errors.New("db error")),
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			// Reset expected calls so they don't accumulate
			suite.photosRepositoryMock.ExpectedCalls = nil
			suite.readingsRepositoryMock.ExpectedCalls = nil
			suite.photosRepositoryMock.
				On("Delete", mock.Anything).
				Return(test.mockReturn[0])
			suite.readingsRepositoryMock.
				On("Delete", mock.Anything).
				Return(test.mockReturn[1])

			err := suite.readingsService.DeleteReading(test.args[0], test.args[1])

			assert.Equal(t, test.want, err)
		})
	}
}

func TestReadingsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ReadingsServiceTestSuite))
}
