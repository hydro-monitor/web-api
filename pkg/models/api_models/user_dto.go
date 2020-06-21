package api_models

type UserDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Password string `json:"password,omitempty"`
	Admin    bool   `json:"admin"`
}
