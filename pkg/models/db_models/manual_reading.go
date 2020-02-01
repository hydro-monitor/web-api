package db_models

var columns = []string{"manual_reading"}

type ManualReadingDTO struct {
	NodeId        string `db:"id"`
	ManualReading bool   `db:"manual_reading"`
}

func (m *ManualReadingDTO) GetColumns() []string {
	return columns
}
