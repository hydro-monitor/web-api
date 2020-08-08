package services

import (
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/models/db_models"
	"hydro_monitor/web_api/pkg/repositories"
	"net/http"
)

type UsersService interface {
	ValidateCredentials(email, password string) (*api_models.UserDTO, ServiceError)
	Register(user *api_models.UserDTO) ServiceError
	GetUserInfo(email string) (*api_models.UserDTO, ServiceError)
	UpdateUser(user *api_models.UserDTO) ServiceError
	DeleteUser(email string) ServiceError
}

type usersServiceImpl struct {
	usersRepository repositories.Repository
}

func (u *usersServiceImpl) UpdateUser(user *api_models.UserDTO) ServiceError {
	dbUser := &db_models.UserDTO{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		Admin:    user.Admin,
	}
	if user.Password != nil {
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		if err != nil {
			return NewGenericServiceError("Error when trying to update user information", err)
		}
		dbUser.Password = encryptedPassword
	}
	dbUser.DetectColumns()
	applied, err2 := u.usersRepository.SafeUpdate(dbUser)
	if !applied {
		return NewNotFoundError("Can't update, user does not exist", err2)
	} else if err2 != nil {
		if err2 == gocql.ErrNotFound {
			return NewNotFoundError("Can't update, user does not exist", err2)
		}
		return NewGenericServiceError("Error when trying to update user information", err2)
	}
	return nil
}

func (u *usersServiceImpl) DeleteUser(email string) ServiceError {
	dbUser := &db_models.UserDTO{Email: &email}
	if err := u.usersRepository.Delete(dbUser); err != nil {
		return NewGenericServiceError("Error when trying to delete user", err)
	}
	return nil
}

func (u *usersServiceImpl) GetUserInfo(email string) (*api_models.UserDTO, ServiceError) {
	dbUser := &db_models.UserDTO{Email: &email}
	if err := u.usersRepository.Get(dbUser); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewNotFoundError("User not found", err)
		}
		return nil, NewGenericServiceError("Error when trying to get user information", err)
	}
	apiUser := &api_models.UserDTO{
		Email:    dbUser.Email,
		Name:     dbUser.Name,
		LastName: dbUser.LastName,
		Admin:    dbUser.Admin,
	}
	return apiUser, nil
}

func (u *usersServiceImpl) Register(user *api_models.UserDTO) ServiceError {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewUserRegistrationError("Error when trying to register a new user", err)
	}
	admin := false
	if user.Admin != nil {
		admin = *user.Admin
	}
	dbUSer := &db_models.UserDTO{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		Password: encryptedPassword,
		Admin:    &admin,
	}
	dbUSer.SetColumns(u.usersRepository.GetColumns())
	if applied, err2 := u.usersRepository.SafeInsert(dbUSer); !applied || err2 != nil {
		if !applied {
			return NewServiceError(http.StatusUnprocessableEntity, "User already exists", err2)
		}
		return NewGenericServiceError("Error when trying to register a new user", err2)
	}
	return nil
}

func (u *usersServiceImpl) ValidateCredentials(email, password string) (*api_models.UserDTO, ServiceError) {
	user := &db_models.UserDTO{Email: &email}
	if err := u.usersRepository.Get(user); err != nil {
		if err == gocql.ErrNotFound {
			return nil, NewInvalidCredentialsError("Wrong username or password", nil)
		}
		return nil, NewGenericServiceError("Error when trying to login", err)
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return nil, NewInvalidCredentialsError("Wrong username or password", err)
	}
	return &api_models.UserDTO{Email: &email, Admin: user.Admin}, nil
}

func NewUsersService(usersRepository repositories.Repository) UsersService {
	return &usersServiceImpl{usersRepository: usersRepository}
}
