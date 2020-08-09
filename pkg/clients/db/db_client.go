package db

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"hydro_monitor/web_api/pkg/models/db_models"
	"time"
)

type Client interface {
	CreateKeyspace(keyspaceName string, replicationFactor int) error
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
	session *gocqlx.Session
}

func (db *clientImpl) CreateKeyspace(keyspaceName string, replicationFactor int) error {
	return db.session.ExecStmt(fmt.Sprintf(
		`CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION ={ 'class' : 'SimpleStrategy', 'replication_factor' : %d }`,
		keyspaceName,
		replicationFactor))
}

func (db *clientImpl) SelectAll(table *table.Table, args db_models.SelectDTO) error {
	stmt, names := qb.Select(table.Name()).ToCql()
	q := db.session.Query(stmt, names)
	return q.SelectRelease(args.GetArgs())
}

func (db *clientImpl) Delete(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Delete(args.GetColumns()...)
	q := db.session.Query(stmt, names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) Select(table *table.Table, args db_models.SelectDTO, pageState []byte, pageSize int) error {
	stmt, names := table.Select(args.GetColumns()...)
	q := db.session.Query(stmt, names).BindMap(args.GetBindMap())
	if pageSize > 0 {
		q.PageState(pageState)
		q.PageSize(pageSize)
	}
	return q.SelectRelease(args.GetArgs())
}

func (db *clientImpl) Update(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Update(args.GetColumns()...)
	q := db.session.Query(stmt, names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) SafeUpdate(table *table.Table, args db_models.DbDTO) (bool, error) {
	stmt, names := table.Update(args.GetColumns()...)
	q := db.session.Query(stmt, names).BindStruct(args)
	return true, q.ExecRelease()
}

func (db *clientImpl) Insert(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Insert()
	q := db.session.Query(stmt, names).BindStruct(args)
	return q.ExecRelease()
}

func (db *clientImpl) SafeInsert(table *table.Table, args db_models.DbDTO) (bool, error) {
	stmt, names := qb.Insert(table.Name()).Columns(args.GetColumns()...).Unique().ToCql()
	q := db.session.Query(stmt, names).BindStruct(args)
	return q.ScanCAS()
}

func NewDB(hosts []string, keyspace string) Client {
	retries := 5
	baseTimeout := time.Second * 10
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.ConnectTimeout = time.Second * 10
	cluster.Keyspace = keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
	cluster.Consistency = gocql.One
	for i := 0; i < retries; i++ {
		session, err := gocqlx.WrapSession(cluster.CreateSession())
		if err != nil {
			log.Error("Hosts: ", hosts)
			log.Error("Failed to connect to database: ", err)
			time.Sleep(baseTimeout)
			baseTimeout *= 2
		} else {
			return &clientImpl{session: &session}
		}
	}
	log.Fatal("Retries exhausted, couldn't connect to database")
	return nil
}

func (db *clientImpl) Migrate(dir string) {
	ctx := context.Background()
	if err := migrate.Migrate(ctx, *db.session, dir); err != nil {
		log.Fatal("Failed to execute database migrations: ", err)
	}
}

func (db *clientImpl) Get(table *table.Table, args db_models.DbDTO) error {
	stmt, names := table.Get(args.GetColumns()...)
	q := db.session.Query(stmt, names).BindStruct(args)
	return q.GetRelease(args)
}

func (db *clientImpl) Close() {
	db.session.Close()
}
