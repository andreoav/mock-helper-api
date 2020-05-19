package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber"
	// "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_projectHttp "github.com/andreoav/mock-helper-api/pkg/project/adapter/http"
	_projectRepo "github.com/andreoav/mock-helper-api/pkg/project/repository/mongo"
	_projectService "github.com/andreoav/mock-helper-api/pkg/project/service"

	_mockHttp "github.com/andreoav/mock-helper-api/pkg/mock/adapter/http"
	_mockService "github.com/andreoav/mock-helper-api/pkg/mock/service"
)

// DatabaseConfig config options
type DatabaseConfig struct {
	URI  string
	Name string
}

// Config struct
type Config struct {
	Database DatabaseConfig
}

// Application struct
type Application struct {
	config Config
}

// NewApplication receives the config
// and returns a pointer to an application
func NewApplication(config Config) *Application {
	return &Application{config}
}

// Start the http server
func (a *Application) Start() {
	app := fiber.New()
	// logger := logrus.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(a.config.Database.URI))

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	defer cancel()

	database := client.Database(a.config.Database.Name)

	projectRepo := _projectRepo.NewProjectRepository(database)
	projectService := _projectService.NewProjectService(projectRepo)
	_projectHttp.NewFiberProjectHandler(app, projectService)

	mockService := _mockService.NewMockService(projectRepo)
	_mockHttp.NewFiberMockHandler(app, mockService)

	if err := app.Listen(9000); err != nil {
		fmt.Println("Shutting down server")
	}
}
