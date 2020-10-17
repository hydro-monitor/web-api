package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models/api_models"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
	"time"
)

type UsersController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type usersControllerImpl struct {
	usersService services.UsersService
}

// UpdateUser godoc
// @Summary Actualiza la información de un usuario
// @Tags users
// @Accept json
// @Param email path string true "Dirección de correo del usuario"
// @Param user body api_models.UserDTO true "Información del usuario"
// @Success 204 "No Content"
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{email} [put]
func (u *usersControllerImpl) UpdateUser(c echo.Context) error {
	email := c.Param("email")
	apiUser := &api_models.UserDTO{Email: &email}
	if err := c.Bind(apiUser); err != nil {
		return echo.ErrBadRequest
	}
	if err := u.usersService.UpdateUser(apiUser); err != nil {
		return err.ToHTTPError()
	}
	return c.NoContent(http.StatusNoContent)
}

// DeleteUser godoc
// @Summary Borra un usuario
// @Tags users
// @Param email path string true "Dirección de correo del usuario"
// @Success 204 "No Content"
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{email} [delete]
func (u *usersControllerImpl) DeleteUser(c echo.Context) error {
	email := c.Param("email")
	if err := u.usersService.DeleteUser(email); err != nil {
		return err.ToHTTPError()
	}
	return c.NoContent(http.StatusNoContent)
}

// GetUser godoc
// @Summary Obtiene la información de un usuario
// @Tags users
// @Produce json
// @Param email path string true "Dirección de correo del usuario"
// @Success 200 {object} api_models.UserDTO
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users/{email} [get]
func (u *usersControllerImpl) GetUser(c echo.Context) error {
	email := c.Param("email")
	user, err := u.usersService.GetUserInfo(email)
	if err != nil {
		return err.ToHTTPError()
	}
	return c.JSON(http.StatusOK, user)
}

// Register godoc
// @Summary Crea un usuario
// @Tags users
// @Accept json
// @Param user body api_models.UserDTO true "Datos del usario"
// @Success 204 "No Content"
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /users [post]
func (u *usersControllerImpl) Register(c echo.Context) error {
	user := new(api_models.UserDTO)
	if err := c.Bind(user); err != nil {
		return echo.ErrBadRequest
	}
	if err := u.usersService.Register(user); err != nil {
		return err.ToHTTPError()
	}
	return c.NoContent(http.StatusCreated)
}

// Login godoc
// @Summary Crea una nueva sesión
// @Tags users
// @Produce json
// @Param username formData string true "Dirección de correo del usuario"
// @Param password formData string true "Contraseña del usuario"
// @Success 200 {object} map[string]string "Token"
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /session [post]
func (u *usersControllerImpl) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userInfo, err := u.usersService.ValidateCredentials(username, password)
	if err != nil {
		return err.ToHTTPError()
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userInfo.Email
	claims["admin"] = userInfo.Admin
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	t, err2 := token.SignedString([]byte("hydromon2020"))
	if err2 != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func NewUsersController(userService services.UsersService) UsersController {
	return &usersControllerImpl{usersService: userService}
}
