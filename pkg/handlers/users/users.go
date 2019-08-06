package users

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
)

func GetUsers(c echo.Context) error {
	users := make([]models.User, 2)
	users[0] = models.User{Email: "example@mail.com", Password: "12245", Admin: false}
	users[1] = models.User{Email: "ryan@air.com", Password: "12245", Admin: true}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, models.User{Email: email, Password: "12245", Admin: false})
}

func CreateUser(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, email)
}
