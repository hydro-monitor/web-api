package db_models

import "github.com/gocql/gocql"

type Reading struct {
	NodeId      string
	ReadingTime gocql.UUID
	WaterLevel  float64
}

func (r *Reading) GetColumns() []string {
	return nil
}
