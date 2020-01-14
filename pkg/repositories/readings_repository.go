package repositories

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

type Reading struct {
	NodeId      string
	ReadingTime gocql.UUID
	WaterLevel  float32
}

func NewReadingsRepository(client db.Client) Repository {
	readingsMetadata := table.Metadata{
		Name:    "readings",
		Columns: []string{"node_id", "reading_time", "water_level"},
		PartKey: []string{"node_id", "reading_time"},
		SortKey: nil,
	}
	return &repositoryImpl{
		table:  table.New(readingsMetadata),
		client: client,
	}
}
