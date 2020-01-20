package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "hydro_monitor/web_api/docs" // docs is generated by Swag CLI, you have to import it.
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/controllers"
	"hydro_monitor/web_api/pkg/services"
	"os"
	"strings"
)

/*import (
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
	e.GET("/nodos/:id/medicion-manual", manual_readings.GetManualReadingFromNode)
	e.PUT("/nodos/:id/medicion-manual", manual_readings.UpdateManualReading)

	// Configurations
	e.GET("/nodos/:id/configuracion", configurations.GetNodeConfiguration)
	e.PUT("/nodos/:id/configuracion", configurations.PutNodeConfiguration)

	// Login route
	e.POST("/login", auth.Login)

	// Unauthenticated route
	e.GET("/", auth.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", auth.Restricted)

	e.Logger.Fatal(e.Start(":1323"))
}*/

// @title Hydro Monitor Web API
// @version 1.0
// @description This is the Hydro Monitor Web API

// @contact.name Manuel Porto
// @contact.email manu.porto94@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api
func main() {
	// Database
	client := db.NewDB(strings.Split(os.Getenv("DB_HOSTS"), ","), os.Getenv("DB_KEYSPACE"))
	client.Migrate("./scripts")
	defer client.Close()

	// Router
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		e.Logger.Fatal("$PORT not set")
		return
	}

	// Services
	nodeService := services.NewNodeService(client)
	readingsService := services.NewReadingsService(client)

	// Controllers
	nodeController := controllers.NewNodeController(nodeService)
	readingsController := controllers.NewReadingsController(readingsService)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	apiGroup := e.Group("/api")

	// Documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Nodes
	nodeGroup := apiGroup.Group("/nodes")
	nodeGroup.GET("/:node_id", nodeController.GetNodeByID).Name = "get-node"
	nodeGroup.GET("/:node_id/configuration", nodeController.GetNodeConfiguration).Name = "get-node-configuration"
	nodeGroup.POST("/:node_id/readings", readingsController.CreateReading)
	nodeGroup.GET("/:node_id/readings/:reading_id", readingsController.GetNodeReading)

	apiGroup.GET("/readings/:reading_id/photos", readingsController.GetReadingPhoto)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
