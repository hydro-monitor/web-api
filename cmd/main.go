package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "hydro_monitor/web_api/docs" // docs is generated by Swag CLI, you have to import it.
	"hydro_monitor/web_api/pkg/clients/db"
	"hydro_monitor/web_api/pkg/controllers"
	"hydro_monitor/web_api/pkg/middlewares"
	"hydro_monitor/web_api/pkg/repositories"
	"hydro_monitor/web_api/pkg/router"
	"hydro_monitor/web_api/pkg/services"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func bootstrapDatabase(keyspaceName string, replicationFactor int) {
	client := db.NewDB(strings.Split(os.Getenv("DB_HOSTS"), ","), "system")
	defer client.Close()
	if err := client.CreateKeyspace(keyspaceName, replicationFactor); err != nil {
		log.Fatal("Couldn't create keyspace")
	}
}

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
	hosts := os.Getenv("DB_HOSTS")
	keyspaceName := os.Getenv("DB_KEYSPACE")
	createKeyspace, err := strconv.ParseBool(os.Getenv("DB_CREATE_KEYSPACE"))
	if err != nil {
		log.Fatal("DB_CREATE_KEYSPACE must be a boolean")
	}

	if createKeyspace {
		replicationFactor, err := strconv.Atoi(os.Getenv("DB_REPLICATION_FACTOR"))
		if err != nil {
			log.Fatal("DB_REPLICATION_FACTOR must be a number")
		}
		bootstrapDatabase(keyspaceName, replicationFactor)
	}

	// Database
	client := db.NewDB(strings.Split(hosts, ","), keyspaceName)
	defer client.Close()

	runMigrations, err := strconv.ParseBool(os.Getenv("DB_RUN_MIGRATIONS"))
	if err != nil {
		log.Fatal("DB_RUN_MIGRATIONS must be a boolean")
	}
	if runMigrations {
		client.Migrate("./scripts/db/migrations")
	}

	// Router
	e := echo.New()
	port := os.Getenv("PORT")
	if port == "" {
		e.Logger.Fatal("$PORT not set")
		return
	}
	s := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		ReadTimeout: 25 * time.Minute,
	}

	// Repositories
	configurationsRepository := repositories.NewConfigurationsRepository(client)
	nodesRepository := repositories.NewNodeRepository(client)
	photosRepository := repositories.NewPhotosRepository(client)
	readingsRepository := repositories.NewReadingsRepository(client)
	usersRepository := repositories.NewUsersRepository(client)

	// Services
	nodeService := services.NewNodeService(configurationsRepository, nodesRepository)
	readingsService := services.NewReadingsService(nodesRepository, photosRepository, readingsRepository)
	usersService := services.NewUsersService(usersRepository)

	// Controllers
	healthController := controllers.NewHealthController()
	nodeController := controllers.NewNodeController(nodeService)
	readingsController := controllers.NewReadingsController(nodeService, readingsService)
	usersController := controllers.NewUsersController(usersService)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    middlewares.JWTSkipper(router.AllowedRequests),
		SigningKey: []byte("hydromon2020"),
	}))
	e.Use(middlewares.AuthorizationMiddleware(
		middlewares.AuthorizationConfig{
			Skipper:   middlewares.AuthorizationSkipper(router.RestrictedUserRequests),
			Validator: middlewares.IsAdmin,
			Service:   usersService,
		}),
	)
	e.Use(middlewares.AuthorizationMiddleware(
		middlewares.AuthorizationConfig{
			Skipper:   middlewares.AuthorizationSkipper(router.RestrictedNodeRequests),
			Validator: middlewares.IsNode,
			Service:   nodeService,
		}),
	)

	// Routes

	// Health
	e.GET("/health", healthController.GetHealthStatus)

	// Documentation
	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	apiGroup := e.Group("/api")

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
	nodeGroup.PUT("/:node_id", nodeController.UpdateNode)
	nodeGroup.PUT("/:node_id/manual-reading", nodeController.UpdateNodeManualReading)
	nodeGroup.POST("/:node_id/readings", readingsController.CreateReading)
	nodeGroup.GET("/:node_id/readings/:reading_id", readingsController.GetNodeReading)
	nodeGroup.DELETE("/:node_id/readings/:reading_id", readingsController.DeleteReading)
	nodeGroup.GET("/:node_id/readings", readingsController.GetNodeReadings)
	nodeGroup.GET("/:node_id/readings/:reading_id/photos", readingsController.GetReadingPhoto)
	nodeGroup.POST("/:node_id/readings/:reading_id/photos", readingsController.AddPhotoToReading)

	// Users
	apiGroup.POST("/session", usersController.Login)
	apiGroup.POST("/users", usersController.Register)
	apiGroup.GET("/users/:email", usersController.GetUser)
	apiGroup.PUT("/users/:email", usersController.UpdateUser)
	apiGroup.DELETE("/users/:email", usersController.DeleteUser)

	e.Logger.Fatal(e.StartServer(s))
}
