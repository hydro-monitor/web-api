package db

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"hydro_monitor/web_api/pkg/models/db_models"
)

type Client interface {
	Migrate(dir string)
	Delete(table *table.Table, args db_models.DbDTO) error
	Get(table *table.Table, args db_models.DbDTO) error
	Insert(table *table.Table, args db_models.DbDTO) error
	SafeInsert(table *table.Table, args db_models.DbDTO) (bool, error)
	Update(table *table.Table, args db_models.DbDTO) error
	SafeUpdate(table *table.Table, args db_models.DbDTO) (bool, error)
	Select(table *table.Table, args db_models.SelectDTO, pageState []byte, pageSize int) error
	SelectAll(table *table.Table, args db_models.SelectDTO) error
	Close()
}

type clientImpl struct {
	session *gocql.Session
}

func (db *clientImpl) SelectAll(table *table.Table, args db_models.SelectDTO) error {
	stmt, names := qb.Select(table.Name()).ToCql()
	q := gocqlx.Query(db.session.Query(stmt), names)
	return q.SelectRelease(args.GetArgs())
}

func (db *clientImpl) Delete(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Delete(args.GetColumns()...)
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) Select(table *table.Table, args db_models.SelectDTO, pageState []byte, pageSize int) error {
	stmt, names := table.Select(args.GetColumns()...)
	q := gocqlx.Query(db.session.Query(stmt), names).BindMap(args.GetBindMap())
	if pageSize > 0 {
		q.PageState(pageState)
		q.PageSize(pageSize)
	}
	return q.SelectRelease(args.GetArgs())
}

func (db *clientImpl) Update(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Update(args.GetColumns()...)
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) SafeUpdate(table *table.Table, args db_models.DbDTO) (bool, error) {
	stmt, names := table.Update(args.GetColumns()...)
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return true, q.ExecRelease()
}

func (db *clientImpl) Insert(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Insert()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) SafeInsert(table *table.Table, args db_models.DbDTO) (bool, error) {
	stmt, names := qb.Insert(table.Name()).Columns(args.GetColumns()...).Unique().ToCql()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.ScanCAS()
}

func NewDB(hosts []string, keyspace string) Client {
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.Keyspace = keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	if err != nil {
		log.Error("Hosts: ", hosts)
		log.Fatal("Failed to connect to database: ", err)
	}
	return &clientImpl{session: session}
}

func (db *clientImpl) Migrate(dir string) {
	ctx := context.Background()
	if err := migrate.Migrate(ctx, db.session, dir); err != nil {
		log.Fatal("Failed to execute database migrations: ", err)
	}
}

func (db *clientImpl) Get(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Get(args.GetColumns()...)
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.GetRelease(args)
}

func (db *clientImpl) Close() {
	db.session.Close()
}
