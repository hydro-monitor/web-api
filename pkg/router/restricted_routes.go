package router

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/middlewares"
)

// AllowedRequests define todos los requests que no requieren ningún tipo de autenticación
var AllowedRequests = []middlewares.Request{
	{Method: echo.GET, Path: "/health"},
	{Method: echo.POST, Path: "/api/session"},
	{Method: echo.POST, Path: "/api/users"},
}

// RestrictedUserRequests define todos los requests que requieren que el usuario tenga permisos de administrador
var RestrictedUserRequests = []middlewares.Request{
	{Method: echo.POST, Path: "/api/nodes"},
	{Method: echo.DELETE, Path: "/api/nodes/:node_id"},
	{Method: echo.PUT, Path: "/api/nodes/:node_id"},
	{Method: echo.POST, Path: "/api/nodes/:node_id/configuration"},
	{Method: echo.PUT, Path: "/api/nodes/:node_id/manual-reading"},
	{Method: echo.DELETE, Path: "/api/nodes/:node_id/readings/:reading_id"},
}

// RestrictedNodeRequests define todos los requests que requieren que el nodo incluya en el JWT su contraseña
var RestrictedNodeRequests = []middlewares.Request{
	{Method: echo.POST, Path: "/api/nodes/:node_id/readings"},
	{Method: echo.POST, Path: "/api/nodes/:node_id/readings/:reading_id/photos"},
}
