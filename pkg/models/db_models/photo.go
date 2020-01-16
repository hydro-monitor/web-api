package db_models

import "github.com/gocql/gocql"

type Photo struct {
	ReadingTime gocql.UUID
	Number      int
	Picture     []byte
}
