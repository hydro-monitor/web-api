package db_models

type UserDTO struct {
	Email    string
	Name     string
	LastName string
	Password []byte
	Admin    bool
	Columns  []string `db:"-"`
}

func (u *UserDTO) GetColumns() []string {
	return u.Columns
}

func (u *UserDTO) SetColumns(columns []string) {
	u.Columns = columns
}

func (u *UserDTO) DetectColumns() {
	u.Columns = make([]string, 0)
	if u.Name != "" {
		u.Columns = append(u.Columns, "name")
	}
	if u.LastName != "" {
		u.Columns = append(u.Columns, "last_name")
	}
	if len(u.Password) > 0 {
		u.Columns = append(u.Columns, "password")
	}
}
