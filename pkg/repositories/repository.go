// El paquete repositories contiene la definición de la interfaz para interactuar con la base de datos.
package repositories

import (
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models/db_models"
)

// Repository es la interfaz que permite interactuar con la base de datos
type Repository interface {
	Delete(args db_models.DbDTO) error
	Get(args db_models.DbDTO) error
	Insert(args db_models.DbDTO) error
	Update(args db_models.DbDTO) error
	Select(args db_models.SelectDTO, pageState []byte, pageSize int) error
	SelectAll(args db_models.SelectDTO) error
}

type repositoryImpl struct {
	table  *table.Table
	client db.Client
}

// Delete borra un registro de la base de datos
func (r *repositoryImpl) Delete(args db_models.DbDTO) error {
	return r.client.Delete(r.table, args)
}

// Get obtiene un registro de la base de datos
func (r *repositoryImpl) Get(args db_models.DbDTO) error {
	return r.client.Get(r.table, args)
}

// Insert crea un registro en la base de datos
func (r repositoryImpl) Insert(args db_models.DbDTO) error {
	return r.client.Insert(r.table, args)
}

// Update actualiza un registro en la base de datos
func (r *repositoryImpl) Update(args db_models.DbDTO) error {
	return r.client.Update(r.table, args)
}

func (r *repositoryImpl) Select(args db_models.SelectDTO, pageState []byte, pageSize int) error {
	return r.client.Select(r.table, args, pageState, pageSize)
}

func (r *repositoryImpl) SelectAll(args db_models.SelectDTO) error {
	return r.client.SelectAll(r.table, args)
}
