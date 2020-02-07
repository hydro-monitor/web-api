package services

import (
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/mocks"
	"testing"
)

type NodeServiceTestSuite struct {
	suite.Suite
	nodeService NodeService
}

func (suite *NodeServiceTestSuite) SetupTest() {
	dbClient := new(mocks.Client)
	suite.nodeService = NewNodeService(dbClient)
}

func TestNodeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NodeServiceTestSuite))
}
