package main

import server "github.com/andreoav/mock-helper-api/pkg/mock-server"

func main() {
	config := server.Config{
		Database: server.DatabaseConfig{
			Name: "mock-server",
			URI:  "mongodb://admin:admin@localhost:27017",
		},
	}

	app := server.NewApplication(config)
	app.Start()
}
