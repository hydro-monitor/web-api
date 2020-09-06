package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Request struct {
	Method string
	Path   string
}

type ValidatorFunc func(echo.Context) bool

type AuthorizationConfig struct {
	Skipper   middleware.Skipper
	Validator ValidatorFunc
}

func AuthorizationSkipper(restrictedPaths []Request) middleware.Skipper {
	return func(c echo.Context) bool {
		request := Request{
			Method: c.Request().Method,
			Path:   c.Path(),
		}
		return !requestInArray(request, restrictedPaths)
	}
}

func AuthorizationMiddleware(config AuthorizationConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultSkipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) || config.Validator(c) {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusForbidden, "Insufficient permissions")
		}
	}
}
