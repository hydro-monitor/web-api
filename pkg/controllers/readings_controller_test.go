package controllers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

/*
	Test suite
*/
type ReadingsControllerTestSuite struct {
	suite.Suite
}

func (suite *ReadingsControllerTestSuite) SetupTest() {

}

func (suite *ReadingsControllerTestSuite) TestCreateReading() {
}

func TestReadingsControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ReadingsControllerTestSuite))
}
