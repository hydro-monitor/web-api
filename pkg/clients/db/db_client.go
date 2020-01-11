package db_client

import (
	"github.com/gocql/gocql"
	"github.com/labstack/gommon/log"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/table"
)

type DbClient interface {
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
		log.Fatal("Failed to connect to database: ", err)
	}
	return &dbClientImpl{session: session}
}
func (db *dbClientImpl) Get(table *table.Table, args interface{}) error {
	stmt, names := table.Get()
	q := gocqlx.Query(db.session.Query(stmt), names).BindStruct(args)
	return q.GetRelease(args)
}

func (db *dbClientImpl) Close() {
	db.session.Close()
}
