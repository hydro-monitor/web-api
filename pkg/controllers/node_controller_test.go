package controllers

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	Test objects
*/
type NodeServiceMock struct {
	mock.Mock
}

func (n *NodeServiceMock) GetNodeConfiguration(nodeId string) (*models.NodeConfiguration, error) {
	args := n.Called(nodeId)
	return args.Get(0).(*models.NodeConfiguration), args.Error(1)
}

/*
	Test suite
*/
type NodeControllerTestSuite struct {
	suite.Suite
	nodeServiceMock *NodeServiceMock
	nodeController  NodeController
	e               *echo.Echo
	rec             *httptest.ResponseRecorder
}

func (suite *NodeControllerTestSuite) SetupTest() {
	suite.nodeServiceMock = new(NodeServiceMock)
	suite.nodeController = NewNodeController(suite.nodeServiceMock)
	suite.e = echo.New()
	suite.rec = httptest.NewRecorder()
}

func (suite *NodeControllerTestSuite) TestGetNodeConfiguration() {
	state1 := models.State{
		Name:        "Normal",
		Interval:    60,
		UpperLimit:  9999999,
		LowerLimit:  -9999999,
		PicturesNum: 0,
		Next:        "Alto",
		Prev:        "Bajo",
	}
	statesMap := make(map[string]*models.State)
	statesMap[state1.Name] = &state1
	expectedNodeConfiguration := models.NodeConfiguration{
		NodeId: "1",
		States: statesMap,
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := suite.e.NewContext(req, suite.rec)
	c.SetPath("/nodes/:node_id/configuration")
	c.SetParamNames("node_id")
	c.SetParamValues("1")

	suite.nodeServiceMock.On("GetNodeConfiguration", "1").Return(&expectedNodeConfiguration, nil)

	_ = suite.nodeController.GetNodeConfiguration(c)
	var response models.NodeConfiguration
	_ = json.Unmarshal(suite.rec.Body.Bytes(), &response)

	assert.Equal(suite.T(), http.StatusOK, suite.rec.Code)
	assert.Equal(suite.T(), expectedNodeConfiguration, response)
}

func TestNodeControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NodeControllerTestSuite))
}
