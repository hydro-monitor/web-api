package repositories

type Repository interface {
	Get(args interface{}) error
	Insert(args interface{}) error
}
