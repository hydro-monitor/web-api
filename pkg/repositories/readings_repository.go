package repositories

import (
	"github.com/scylladb/gocqlx/v2/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

// NewReadingsRepository crea un nuevo repositorio para interactuar con la tabla de readings de la base de datos
func NewReadingsRepository(client db.Client) Repository {
	readingsMetadata := table.Metadata{
		Name:    "readings",
		Columns: []string{"node_id", "reading_id", "reading_time", "water_level", "manual_reading"},
		PartKey: []string{"node_id"},
		SortKey: []string{"reading_id"},
	}
	return &repositoryImpl{
		table:  table.New(readingsMetadata),
		client: client,
	}
}
