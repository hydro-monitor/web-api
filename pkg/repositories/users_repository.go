package repositories

import (
	"github.com/scylladb/gocqlx/v2/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

func NewUsersRepository(client db.Client) Repository {
	usersMetadata := table.Metadata{
		Name:    "users",
		Columns: []string{"email", "name", "last_name", "password", "admin"},
		PartKey: []string{"email"},
		SortKey: nil,
	}
	return &repositoryImpl{
		table:  table.New(usersMetadata),
		client: client,
	}
}
