package controllers

import (
	"github.com/bmizerany/assert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StateControllerTestSuite struct {
	suite.Suite
	stateController StateController
}

func (suite *StateControllerTestSuite) SetupTest() {
	suite.stateController = NewStateController()
}

func (suite *StateControllerTestSuite) TestGetConfiguration() {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/nodes/:node_id/configuration")
	c.SetParamNames("node_id")
	c.SetParamValues("1")

	// Assertions
	if err := suite.stateController.GetConfiguration(c); err == nil {
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
	}
}

func TestStateControllerTestSuite(t *testing.T) {
	suite.Run(t, new(StateControllerTestSuite))
}
