package db_models

type NodeDTO struct {
	Id            string `db:"id"`
	Description   string `db:"description"`
	Configuration string `db:"configuration"`
	State         string `db:"state"`
	ManualReading bool   `db:"manual_reading"`
}

func (n *NodeDTO) GetColumns() []string {
	return nil
}
