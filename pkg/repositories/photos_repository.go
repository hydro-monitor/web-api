package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewPhotosRepository(client db.Client) Repository {
	photosMetadata := table.Metadata{
		Name:    "photos",
		Columns: []string{"reading_time", "number", "picture"},
		PartKey: []string{"reading_time", "number"},
		SortKey: nil,
	}
	return &repositoryImpl{
		table:  table.New(photosMetadata),
		client: client,
	}
}
