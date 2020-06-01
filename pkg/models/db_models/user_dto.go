package db_models

type UserDTO struct {
	Email    string
	Password []byte
	Admin    bool
}

func (u *UserDTO) GetColumns() []string {
	return nil
}
