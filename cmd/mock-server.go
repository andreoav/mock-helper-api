package main

import server "github.com/andreoav/mock-helper-api/pkg/mock-server"

func main() {
	config := &server.Config{}
	app := server.NewApplication(config)
	app.Start()
}
