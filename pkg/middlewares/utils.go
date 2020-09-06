package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/services"
)

func IsAdmin(c echo.Context, service interface{}) bool {
	userService, ok := service.(services.UsersService)
	if !ok {
		c.Logger().Error("Wrong user service provided, can not check info provided in JWT")
		return false
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	mail, ok := claims["username"].(string)
	if !ok {
		return false
	}
	apiUser, err := userService.GetUserInfo(mail)
	if err != nil {
		c.Logger().Error("Error when trying to retrieve user's info", err)
		return false
	}
	return *apiUser.Admin
}

func IsNode(c echo.Context, service interface{}) bool {
	nodeService, ok := service.(services.NodeService)
	if !ok {
		c.Logger().Error("Wrong node service provided, can not check info provided in JWT")
		return false
	}
	nodeId := c.Param("node_id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	password, ok := claims["password"].(string)
	if !ok {
		return false
	}
	valid, err := nodeService.CheckNodeCredentials(nodeId, password)
	if err != nil {
		c.Logger().Error("Error when trying to check node credentials", err)
	}
	return valid
}

func requestInArray(request Request, requests []Request) bool {
	for _, r := range requests {
		if r.Path == request.Path && r.Method == request.Method {
			return true
		}
	}
	return false
}
