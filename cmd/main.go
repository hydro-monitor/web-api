package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "hydro_monitor/web_api/docs" // docs is generated by Swag CLI, you have to import it.
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/controllers"
	"hydro_monitor/web_api/pkg/repositories"
	"hydro_monitor/web_api/pkg/services"
	"net/http"
	"os"
	"strings"
	"time"
)

// @title Hydro Monitor Web API
// @version 0.1.0
// @description Esta es la definición de la API del servidor del Hydro Monitor

// @contact.name Agustina Barbetta
// @contact.email agustina.barbetta@gmail.com

// @contact.name Manuel Porto
// @contact.email manu.porto94@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

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
	s := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		ReadTimeout: 5 * time.Minute,
	}

	// Repositories
	configurationsRepository := repositories.NewConfigurationsRepository(client)
	nodesRepository := repositories.NewNodeRepository(client)
	photosRepository := repositories.NewPhotosRepository(client)
	readingsRepository := repositories.NewReadingsRepository(client)

	// Services
	nodeService := services.NewNodeService(configurationsRepository, nodesRepository)
	readingsService := services.NewReadingsService(nodesRepository, photosRepository, readingsRepository)

	// Controllers
	nodeController := controllers.NewNodeController(nodeService)
	readingsController := controllers.NewReadingsController(nodeService, readingsService)

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
	nodeGroup.GET("/last-reading", readingsController.GetNodesLastReading)
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

	e.Logger.Fatal(e.StartServer(s))
}
