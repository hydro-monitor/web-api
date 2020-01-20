package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
)

type Repository interface {
	Get(args interface{}) error
	Insert(args interface{}) error
}

type repositoryImpl struct {
	table  *table.Table
	client db.Client
}

func (r repositoryImpl) Get(args interface{}) error {
	return r.client.Get(r.table, args)
}

func (r repositoryImpl) Insert(args interface{}) error {
	return r.client.Insert(r.table, args)
}
