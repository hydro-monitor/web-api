package services

import (
	"bytes"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
)

type UsersService interface {
	ValidateCredentials(email, password string) (*api_models.UserDTO, ServiceError)
	Register(user *api_models.UserDTO) ServiceError
}

type usersServiceImpl struct {
	usersRepository repositories.Repository
}

func (u *usersServiceImpl) Register(user *api_models.UserDTO) ServiceError {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewUserRegistrationError("Error when trying to register a new user", err)
	}
	dbUSer := &db_models.UserDTO{
		Email:    user.Email,
		Password: encryptedPassword,
		Admin:    user.Admin,
	}
	if err2 := u.usersRepository.Insert(dbUSer); err2 != nil {
		return NewGenericServiceError("Error when trying to register a new user", err2)
	}
	return nil
}

func (u *usersServiceImpl) ValidateCredentials(email, password string) (*api_models.UserDTO, ServiceError) {
	user := &db_models.UserDTO{Email: email}
	if err := u.usersRepository.Get(user); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewInvalidCredentialsError("Wrong username or password", nil)
		}
		return nil, NewGenericServiceError("Error when trying to login", err)
	}
	encryptedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err2 != nil {
		return nil, NewGenericServiceError("Error when trying to login", err2)
	}
	if !bytes.Equal(user.Password, encryptedPassword) {
		return nil, NewInvalidCredentialsError("Wrong username or password", nil)
	}
	return &api_models.UserDTO{Email: email, Admin: true}, nil
}

func NewUsersService(usersRepository repositories.Repository) UsersService {
	return &usersServiceImpl{usersRepository: usersRepository}
}
