package repositories

import (
	"github.com/scylladb/gocqlx/v2/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

// NewNodesRepository crea un nuevo repositorio para interactuar con la tabla de "nodes" de la base de datos
func NewNodeRepository(client db.Client) Repository {
	nodeMetadata := table.Metadata{
		Name:    "nodes",
		Columns: []string{"id", "description", "manual_reading"},
		PartKey: []string{"id"},
		SortKey: nil,
	}
	return &repositoryImpl{table: table.New(nodeMetadata), client: client}
}
