package db_driver

import "github.com/gocql/gocql"

type Driver struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func NewDriver() *Driver {
	// TODO get this info from an env file or similar
	cluster := gocql.NewCluster("192.168.50.41", "192.168.50.42", "192.168.50.43")
	cluster.Keyspace = "usertest"
	cluster.PoolConfig.HostSelectionPolicy = gocql.RoundRobinHostPolicy()
	return &Driver{cluster: cluster}
}

func (d *Driver) StartSession() error {
	session, err := d.cluster.CreateSession()
	d.session = session
	return err
}

func (d *Driver) EndSession() {
	d.session.Close()
}
