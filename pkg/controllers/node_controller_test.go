package controllers

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/mocks"
	"hydro_monitor/web_api/pkg/models/api_models"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	Test suite
*/
type NodeControllerTestSuite struct {
	suite.Suite
	nodeServiceMock *mocks.NodeService
	nodeController  NodeController
	e               *echo.Echo
	rec             *httptest.ResponseRecorder
}

func (suite *NodeControllerTestSuite) SetupTest() {
	suite.nodeServiceMock = new(mocks.NodeService)
	suite.nodeController = NewNodeController(suite.nodeServiceMock)
	suite.e = echo.New()
	suite.rec = httptest.NewRecorder()
}

func (suite *NodeControllerTestSuite) TestGetNode() {
	id := "1"
	description := "A node"
	node := api_models.NodeDTO{
		Id:          &id,
		Description: &description,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := suite.e.NewContext(req, suite.rec)
	c.SetPath("/nodes/:node_id")
	c.SetParamNames("node_id")
	c.SetParamValues("1")

	suite.nodeServiceMock.On("GetNode", "1").Return(&node, nil)

	_ = suite.nodeController.GetNodeByID(c)
	var response api_models.NodeDTO
	_ = json.Unmarshal(suite.rec.Body.Bytes(), &response)

	assert.Equal(suite.T(), http.StatusOK, suite.rec.Code)
	assert.Equal(suite.T(), node, response)
}

func TestNodeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NodeControllerTestSuite))
}
