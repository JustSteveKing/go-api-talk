package infrastructure

import (
	"net/http"
	"time"

	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure/factories"
	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure/providers"
	gohandlers "github.com/gorilla/handlers"
)

// Boot our application
func Boot() *factories.Application {
	// Build our config
	config := providers.ConfigProvider()

	// Build our router
	router := providers.RouteProvider()

	// Build our logger
	logger := providers.LoggerProvider()

	// Set up CORS protection
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	return &factories.Application{
		Server: &http.Server{
			Addr:         ":" + config.App.Port,
			Handler:      corsHandler(router),
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Router: router,
		Logger: logger,
		Config: config,
	}
}
