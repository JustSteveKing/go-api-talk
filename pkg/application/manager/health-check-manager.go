package manager

import (
	"net/http"

	"github.com/JustSteveKing/go-api-talk/pkg/application/health/handlers"

	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure/factories"
)

// SetupHealthCheckService will set up all service level requirements for the Health Check service
func SetupHealthCheckService(app *factories.Application) {
	// Create the HTTP Handler
	handler := handlers.NewHandler(app)

	// Create a service Router
	router := app.Router.Methods(http.MethodGet).Subrouter()

	// Register our service routes
	router.HandleFunc("/ping", handler.Handle).Name("health:ping")
}
