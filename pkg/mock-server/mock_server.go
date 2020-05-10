package server

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_projectHttp "github.com/andreoav/mock-helper-api/pkg/project/adapter/http"
	_projectRepo "github.com/andreoav/mock-helper-api/pkg/project/repository/mongo"
	_projectService "github.com/andreoav/mock-helper-api/pkg/project/service"
)

type Config struct{}

type Application struct {
	config *Config
}

func NewApplication(config *Config) *Application {
	return &Application{config}
}

func (a *Application) Start() {
	app := fiber.New()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin@localhost:27017"))

	if err != nil {
		panic("Application was not able to start")
	}

	database := client.Database("mock-server")

	projectRepo := _projectRepo.NewProjectRepository(database)
	projectService := _projectService.NewProjectService(projectRepo)
	_projectHttp.NewFiberProjectHandler(app, projectService)

	log.Fatal(app.Listen(9000))
}
