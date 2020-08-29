// El paquete repositories contiene la definición de la interfaz para interactuar con la base de datos.
package repositories

import (
	"github.com/scylladb/gocqlx/v2/table"
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/models/db_models"
)

// Repository es la interfaz que permite interactuar con la base de datos
type Repository interface {
	GetColumns() []string
	GetPartitionKey() []string
	GetSortingKey() []string
	Delete(args db_models.DbDTO) error
	Get(args db_models.DbDTO) error
	Insert(args db_models.DbDTO) error
	SafeInsert(args db_models.DbDTO) (bool, error)
	Update(args db_models.DbDTO) error
	SafeUpdate(args db_models.DbDTO) (bool, error)
	Select(args db_models.SelectDTO, pageState []byte, pageSize int) ([]byte, error)
	SelectAll(args db_models.SelectDTO) error
}

type repositoryImpl struct {
	table  *table.Table
	client db.Client
}

func (r *repositoryImpl) GetColumns() []string {
	return r.table.Metadata().Columns
}

func (r *repositoryImpl) GetPartitionKey() []string {
	return r.table.Metadata().PartKey
}

func (r *repositoryImpl) GetSortingKey() []string {
	return r.table.Metadata().SortKey
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

func (r repositoryImpl) SafeInsert(args db_models.DbDTO) (bool, error) {
	return r.client.SafeInsert(r.table, args)
}

// Update actualiza un registro en la base de datos
func (r *repositoryImpl) Update(args db_models.DbDTO) error {
	return r.client.Update(r.table, args)
}

func (r *repositoryImpl) SafeUpdate(args db_models.DbDTO) (bool, error) {
	return r.client.SafeUpdate(r.table, args)
}

func (r *repositoryImpl) Select(args db_models.SelectDTO, pageState []byte, pageSize int) ([]byte, error) {
	return r.client.Select(r.table, args, pageState, pageSize)
}

func (r *repositoryImpl) SelectAll(args db_models.SelectDTO) error {
	return r.client.SelectAll(r.table, args)
}
