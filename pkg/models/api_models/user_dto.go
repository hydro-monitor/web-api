package api_models

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Admin    bool   `json:"admin"`
}
