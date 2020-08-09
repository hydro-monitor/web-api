package db_models

import "github.com/scylladb/gocqlx/v2/qb"

type SelectDTO interface {
	GetColumns() []string
	GetBindMap() qb.M
	GetArgs() interface{}
}
