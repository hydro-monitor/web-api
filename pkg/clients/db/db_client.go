package db

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/gocqlx/table"
)

type Client interface {
	Migrate(dir string)
	Get(table *table.Table, args interface{}) error
	Insert(table *table.Table, args interface{}) error
	Close()
}

type clientImpl struct {
	session *gocql.Session
}

func (db *clientImpl) Insert(table *table.Table, args interface{}) error {
	stmt, names := table.Insert()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.ExecRelease()
}

func NewDB(hosts []string, keyspace string) Client {
	cluster := gocql.NewCluster()
	cluster.Hosts = hosts
	cluster.Keyspace = keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
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

func (db *clientImpl) Get(table *table.Table, args interface{}) error {
	stmt, names := table.Get()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.GetRelease(args)
}

func (db *clientImpl) Close() {
	db.session.Close()
}