package db_driver

import (
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Driver struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var driver *Driver

func init() {
	// TODO do this in the program's entry point
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	driver = newDriver()
}

func newDriver() *Driver {
	cluster := gocql.NewCluster()
	cluster.Hosts = strings.Split(os.Getenv("DB_HOSTS"), ",")
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
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
