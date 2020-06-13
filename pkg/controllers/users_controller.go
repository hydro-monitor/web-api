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
}

type usersControllerImpl struct {
	usersService services.UsersService
}

func (u *usersControllerImpl) GetUser(c echo.Context) error {
	mail := c.Param("mail")

}

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

	t, err2 := token.SignedString([]byte("secret"))
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
