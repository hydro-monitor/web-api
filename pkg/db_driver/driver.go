package db_driver

import (
	"github.com/gocql/gocql"
	"log"
)

type Driver struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var driver *Driver

func init() {
	driver = newDriver()
}

func newDriver() *Driver {
	// TODO get this info from an env file or similar
	cluster := gocql.NewCluster("192.168.50.41", "192.168.50.42", "192.168.50.43")
	cluster.Keyspace = "hydromonitor"
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return &Driver{cluster: cluster, session: session}
}

func GetDriver() *Driver {
	return driver
}

func (d *Driver) GetSession() *gocql.Session {
	return d.session
}

func (d *Driver) EndSession() {
	d.session.Close()
}
