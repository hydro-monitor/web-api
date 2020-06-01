package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewUsersRepository(client db.Client) Repository {
	usersMetadata := table.Metadata{
		Name:    "users",
		Columns: []string{"email", "password", "admin"},
		PartKey: []string{"email"},
		SortKey: nil,
	}
	return &repositoryImpl{
		table:  table.New(usersMetadata),
		client: client,
	}
}
