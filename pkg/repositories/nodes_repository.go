package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewNodeRepository(client db.Client) Repository {
	nodeMetadata := table.Metadata{
		Name:    "nodes",
		Columns: []string{"id", "description", "manual_reading"},
		PartKey: []string{"id"},
		SortKey: nil,
	}
	return &repositoryImpl{table: table.New(nodeMetadata), client: client}
}
