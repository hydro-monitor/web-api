package router

import (
	"github.com/labstack/echo/v4"
	"hydro_monitor/web_api/pkg/middlewares"
)

// RestrictedUserRoutes defines all routes that require the user to have an admin role
var RestrictedUserRoutes = []middlewares.Request{
	{Method: echo.POST, Path: "/api/nodes"},
	{Method: echo.DELETE, Path: "/api/nodes/:node_id"},
	{Method: echo.PUT, Path: "/api/nodes/:node_id"},
	{Method: echo.POST, Path: "/api/nodes/:node_id/configuration"},
	{Method: echo.PUT, Path: "/api/nodes/:node_id/manual-reading"},
	{Method: echo.DELETE, Path: "/api/nodes/:node_id/readings/:reading_id"},
}

// RestrictedNodeRoutes defines all routes that require that the node includes it's token
var RestrictedNodeRoutes = []middlewares.Request{
	{Method: echo.POST, Path: "/api/nodes/:node_id/readings"},
	{Method: echo.POST, Path: "/api/nodes/:node_id/readings/:reading_id/photos"},
}
