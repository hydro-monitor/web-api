package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func IsAdmin(c echo.Context) bool {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	admin, ok := claims["admin"].(bool)
	if !ok {
		return false
	}
	return admin
}

func IsNode(c echo.Context) bool {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	token, ok := claims["token"].(string)
	if !ok {
		return false
	}
	if len(token) == 16 {
		return true
	}
	return false
}

func requestInArray(request Request, requests []Request) bool {
	for _, r := range requests {
		if r.Path == request.Path && r.Method == request.Method {
			return true
		}
	}
	return false
}
