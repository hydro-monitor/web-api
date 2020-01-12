package db_client

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/gocqlx/table"
)

type DbClient interface {
	Migrate(dir string)
	Get(table *table.Table, args interface{}) error
	Close()
}

type dbClientImpl struct {
	session *gocql.Session
}

func NewDB(hosts []string, keyspace string) DbClient {
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.Keyspace = keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
	session, err := cluster.CreateSession()
	if err != nil {
		log.Error("Hosts: ", hosts)
		log.Fatal("Failed to connect to database: ", err)
	}
	return &dbClientImpl{session: session}
}

func (db *dbClientImpl) Migrate(dir string) {
	ctx := context.Background()
	if err := migrate.Migrate(ctx, db.session, dir); err != nil {
		log.Fatal("Failed to execute database migrations: ", err)
	}
}

func (db *dbClientImpl) Get(table *table.Table, args interface{}) error {
	stmt, names := table.Get()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.GetRelease(args)
}

func (db *dbClientImpl) Close() {
	db.session.Close()
}
