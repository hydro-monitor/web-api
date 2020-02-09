package db_models

import "github.com/scylladb/gocqlx/qb"

type SelectDTO interface {
	GetColumns() []string
	GetBindMap() qb.M
	GetArgs() interface{}
}
