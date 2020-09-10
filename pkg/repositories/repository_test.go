package repositories

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"hydro_monitor/web_api/pkg/clients/db"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
	dbClient db.Client
}

func (suite *RepositoryTestSuite) SetupTest() {
	suite.dbClient = new(db.ClientMock)
}

func (suite *RepositoryTestSuite) TestNewRepository() {
	tests := []struct {
		name       string
		repository Repository
		columns    []string
		partKey    []string
		sortKey    []string
	}{
		{
			name:       "Configurations repository it's not nil",
			repository: NewConfigurationsRepository(suite.dbClient),
			columns:    []string{"node_id", "configuration"},
			partKey:    []string{"node_id"},
			sortKey:    nil,
		},
		{
			name:       "Nodes repository it's not nil",
			repository: NewNodeRepository(suite.dbClient),
			columns:    []string{"id", "description", "manual_reading", "password"},
			partKey:    []string{"id"},
			sortKey:    nil,
		},
		{
			name:       "Photos repository it's not nil",
			repository: NewPhotosRepository(suite.dbClient),
			columns:    []string{"reading_time", "number", "picture"},
			partKey:    []string{"reading_time", "number"},
			sortKey:    nil,
		},
		{
			name:       "Readings repository it's not nil",
			repository: NewReadingsRepository(suite.dbClient),
			columns:    []string{"node_id", "reading_id", "reading_time", "water_level", "manual_reading"},
			partKey:    []string{"node_id"},
			sortKey:    []string{"reading_id"},
		},
		{
			name:       "Users repository it's not nil",
			repository: NewUsersRepository(suite.dbClient),
			columns:    []string{"email", "name", "last_name", "password", "admin"},
			partKey:    []string{"email"},
			sortKey:    nil,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			assert.NotNil(t, test.repository)
			assert.Equal(t, test.columns, test.repository.GetColumns())
			assert.Equal(t, test.partKey, test.repository.GetPartitionKey())
			assert.Equal(t, test.sortKey, test.repository.GetSortingKey())
		})
	}
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
