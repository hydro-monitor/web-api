package db_models

type DbDTO interface {
	GetColumns() []string
}
