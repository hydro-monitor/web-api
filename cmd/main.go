package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/handlers/auth"
	"hydro_monitor/web_api/pkg/handlers/configurations"
	"hydro_monitor/web_api/pkg/handlers/manual_readings"
	"hydro_monitor/web_api/pkg/handlers/nodes"
	"hydro_monitor/web_api/pkg/handlers/readings"
	"hydro_monitor/web_api/pkg/handlers/users"
)

func main() {
	defer db_driver.GetDriver().EndSession()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Users
	e.GET("/usuarios", users.GetUsers)
	e.GET("/usuarios/:mail", users.GetUser)
	e.POST("/usuarios", users.CreateUser)
	e.DELETE("/usuarios/:mail", users.DeleteUser)

	// Nodes
	e.GET("/nodos", nodes.GetNodes)
	e.GET("/nodos/:id", nodes.GetNode)
	e.POST("/nodos", nodes.PostNode)
	e.DELETE("/nodos/:id", nodes.DeleteNode)

	// Readings
	e.GET("/nodos/:id/mediciones", readings.GetAllNodeReadings)
	e.POST("/nodos/:id/mediciones", readings.PostReading)
	e.DELETE("/nodos/:id/mediciones/:timestamp", readings.DeleteReading)

	// Manual Readings
	e.GET("/nodos/:id/medicion_manual", manual_readings.GetManualReadingFromNode)
	e.PUT("/nodos/:id/medicion_manual", manual_readings.UpdateManualReading)

	// Configurations
	e.GET("/nodos/:id/configuracion/", configurations.GetNodeConfiguration)
	e.POST("nodos/:id/configuracion/", configurations.PostNodeConfiguration)

	// Login route
	e.POST("/login", auth.Login)

	// Unauthenticated route
	e.GET("/", auth.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", auth.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
