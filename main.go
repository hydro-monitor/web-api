package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hydro_monitor/web_api/auth"
	"hydro_monitor/web_api/users"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Users
	e.GET("/usuarios", users.GetUsers)
	e.GET("/usuarios/:mail", users.GetUser)
	e.POST("/usuarios", users.CreateUser)
	e.DELETE("/usuarios/:mail", users.DeleteUser)

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

func Hello() string {
	return "Hello, world."
}