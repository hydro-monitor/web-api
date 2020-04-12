package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewConfigurationsRepository(client db.Client) Repository {
	configurationsMetadata := table.Metadata{
		Name:    "configurations",
		Columns: []string{"node_id", "configuration"},
		PartKey: []string{"node_id"},
		SortKey: nil,
	}
	return &repositoryImpl{table: table.New(configurationsMetadata), client: client}
}
