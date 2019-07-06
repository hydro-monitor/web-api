package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Email string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Admin bool `json:"admin" form:"admin" query:"admin"`
}

func GetUsers(c echo.Context) error {
	users := make([]User, 2)
	users[0] = User{"example@mail.com", "12245", false}
	users[1] = User{"ryan@air.com", "12245", true}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, User{email, "12245", false})
}

func CreateUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	email := c.QueryParam("email")
	return c.JSON(http.StatusOK, email)
}