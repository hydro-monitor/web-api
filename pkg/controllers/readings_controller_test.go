package controllers

import (
	"errors"
	"github.com/bmizerany/assert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	Test suite
*/
type ReadingsControllerTestSuite struct {
	suite.Suite
	nodesServiceMock    *services.NodeServiceMock
	readingsServiceMock *services.ReadingsServiceMock
	readingsController  ReadingsController
	e                   *echo.Echo
	rec                 *httptest.ResponseRecorder
}

func (suite *ReadingsControllerTestSuite) setupTest() {
	suite.nodesServiceMock = new(services.NodeServiceMock)
	suite.readingsServiceMock = new(services.ReadingsServiceMock)
	suite.readingsController = NewReadingsController(suite.nodesServiceMock, suite.readingsServiceMock)
	suite.e = echo.New()
	suite.rec = httptest.NewRecorder()
}

func (suite *ReadingsControllerTestSuite) TestDeleteReading() {
	tests := []struct {
		name         string
		args         []string
		mockResponse services.ServiceError
		want         int
	}{
		{
			name:         "ok",
			args:         []string{"node-ok", "00336270-8191-11ea-a43d-0242ac120003"},
			mockResponse: nil,
			want:         http.StatusNoContent,
		},
		{
			name:         "when node ID it's not provided should return a bad request error",
			args:         []string{"", "00336270-8191-11ea-a43d-0242ac120003"},
			mockResponse: nil,
			want:         http.StatusBadRequest,
		},
		{
			name:         "when reading ID it's not provided should return a bad request error",
			args:         []string{"node-ok", ""},
			mockResponse: nil,
			want:         http.StatusBadRequest,
		},
		{
			name:         "when neither node ID nor reading ID are not provided should return a bad request error",
			args:         []string{"", ""},
			mockResponse: nil,
			want:         http.StatusBadRequest,
		},
		{
			name:         "when there is a service error should return that same error",
			args:         []string{"node-error", "00336270-8191-11ea-a43d-0242ac120003"},
			mockResponse: services.NewGenericServiceError("db error", errors.New("db error")),
			want:         http.StatusInternalServerError,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			suite.setupTest()
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			c := suite.e.NewContext(req, suite.rec)
			c.SetPath("/nodes/:node_id/readings/:reading_id")
			c.SetParamNames("node_id", "reading_id")
			c.SetParamValues(test.args[0], test.args[1])

			suite.readingsServiceMock.On("DeleteReading", test.args[0], test.args[1]).Return(test.mockResponse)

			if err := suite.readingsController.DeleteReading(c); err != nil {
				he, ok := err.(*echo.HTTPError)
				assert.Equal(t, true, ok)
				if ok {
					assert.Equal(t, test.want, he.Code)
				}
			} else {
				assert.Equal(t, test.want, suite.rec.Code)
			}
		})
	}
}

func TestReadingsControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ReadingsControllerTestSuite))
}
