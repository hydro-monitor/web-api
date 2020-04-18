package db_models

type DeleteNodeDTO struct {
	Id string
}

func (n *DeleteNodeDTO) GetColumns() []string {
	return nil
}
