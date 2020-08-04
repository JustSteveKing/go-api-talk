package main

import (
	"github.com/JustSteveKing/go-api-talk/pkg/application/manager"
	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure"
	"github.com/joho/godotenv"
)

func main() {
	// Load in our environment variables
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	// Create our application
	app := infrastructure.Boot()

	// Register our API services
	manager.SetupHealthCheckService(app)

	// Run our application in a goroutine
	go func() {
		app.Run()
	}()

	// Wait for a termination signal
	app.WaitForShutdown()
}
