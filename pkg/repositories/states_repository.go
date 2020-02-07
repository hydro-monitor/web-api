package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewStatesRepository(client db.Client) Repository {
	stateMetadata := table.Metadata{
		Name: "states",
		Columns: []string{
			"node_id",
			"name",
			"photos_per_reading",
			"reading_interval",
			"lower_limit",
			"upper_limit",
			"previous_state",
			"next_state",
		},
		PartKey: []string{"node_id"},
		SortKey: []string{"name"},
	}
	return &repositoryImpl{table: table.New(stateMetadata), client: client}
}
