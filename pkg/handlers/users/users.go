package users

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/models"
	"net/http"
)

func GetUsers(c echo.Context) error {
	users := make([]models.User, 2)
	users[0] = models.User{"example@mail.com", "12245", false}
	users[1] = models.User{"ryan@air.com", "12245", true}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, models.User{email, "12245", false})
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
