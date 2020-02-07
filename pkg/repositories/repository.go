package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models/db_models"
)

type Repository interface {
	Get(args db_models.DbDTO) error
	Insert(args db_models.DbDTO) error
	Update(args db_models.DbDTO) error
	Select(args db_models.SelectDTO) error
}

type repositoryImpl struct {
	table  *table.Table
	client db.Client
}

func (r repositoryImpl) Get(args db_models.DbDTO) error {
	return r.client.Get(r.table, args)
}

func (r repositoryImpl) Insert(args db_models.DbDTO) error {
	return r.client.Insert(r.table, args)
}

func (r *repositoryImpl) Update(args db_models.DbDTO) error {
	return r.client.Update(r.table, args)
}

func (r *repositoryImpl) Select(args db_models.SelectDTO) error {
	return r.client.Select(r.table, args)
}
