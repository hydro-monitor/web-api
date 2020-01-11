package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

type nodeRepositoryImpl struct {
	table    *table.Table
	dbClient db_client.DbClient
}

func NewNodeRepository(dbClient db_client.DbClient) Repository {
	nodeMetadata := table.Metadata{
		Name:    "nodes",
		Columns: []string{"id", "description", "configuration", "state"},
		PartKey: []string{"id"},
		SortKey: nil,
	}
	return &nodeRepositoryImpl{table: table.New(nodeMetadata), dbClient: dbClient}
}

func (n *nodeRepositoryImpl) Get(args interface{}) error {
	return n.dbClient.Get(n.table, args)
}

func (n *nodeRepositoryImpl) Insert(args interface{}) error {
	panic("implement me")
}
