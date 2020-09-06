package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var AllowedRequests = []Request{
	{Method: echo.GET, Path: "/health"},
	{Method: echo.POST, Path: "/api/session"},
	{Method: echo.POST, Path: "/api/users"},
}

func JWTSkipper(allowedRequests []Request) middleware.Skipper {
	return func(c echo.Context) bool {
		request := Request{
			Method: c.Request().Method,
			Path:   c.Path(),
		}
		return requestInArray(request, allowedRequests)
	}
}
