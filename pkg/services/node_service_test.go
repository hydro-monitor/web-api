package services

import (
	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/pkg/models"
	"testing"
)

type NodeServiceTestSuite struct {
	suite.Suite
	nodeService NodeService
}

func (suite *NodeServiceTestSuite) SetupTest() {
	suite.nodeService = NewNodeService()
}

func (suite *NodeServiceTestSuite) TestGetNodeConfiguration() {
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

	nodeConfiguration, _ := suite.nodeService.GetNodeConfiguration("1")

	assert.Equal(suite.T(), expectedNodeConfiguration, *nodeConfiguration)
}

func TestNodeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NodeServiceTestSuite))
}
