package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

// NewPhotosRepository crea un nuevo repositorio para interactuar con la tabla de "photos" de la base de datos
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
