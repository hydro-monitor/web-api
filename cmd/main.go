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
	e.Use(middleware.CORS())

	// Routes
	apiGroup := e.Group("/api")

	// Documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Nodes
	nodeGroup := apiGroup.Group("/nodes")
	nodeGroup.GET("", nodeController.GetNodes)
	nodeGroup.POST("", nodeController.PostNode)
	nodeGroup.GET("/:node_id", nodeController.GetNodeByID)
	nodeGroup.DELETE("/:node_id", nodeController.DeleteNode)
	nodeGroup.GET("/:node_id/configuration", nodeController.GetNodeConfiguration)
	nodeGroup.POST("/:node_id/configuration", nodeController.CreateNodeConfiguration)
	nodeGroup.GET("/:node_id/manual-reading", nodeController.GetNodeManualReadingStatus)
	nodeGroup.PUT("/:node_id/manual-reading", nodeController.UpdateNodeManualReading)
	nodeGroup.POST("/:node_id/readings", readingsController.CreateReading)
	nodeGroup.GET("/:node_id/readings/:reading_id", readingsController.GetNodeReading)
	nodeGroup.GET("/:node_id/readings", readingsController.GetNodeReadings)
	nodeGroup.GET("/:node_id/readings/:reading_id/photos", readingsController.GetReadingPhoto)
	nodeGroup.POST("/:node_id/readings/:reading_id/photos", readingsController.AddPhotoToReading)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
