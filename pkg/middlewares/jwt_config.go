package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTSkipper(allowedRequests []Request) middleware.Skipper {
	return func(c echo.Context) bool {
		c.Logger().Errorf("Authorization: %s", c.Request().Header.Get("Authorization"))
		request := Request{
			Method: c.Request().Method,
			Path:   c.Path(),
		}
		return requestInArray(request, allowedRequests)
	}
}
