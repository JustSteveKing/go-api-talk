package handlers

import (
	"net/http"

	"github.com/JustSteveKing/go-api-talk/pkg/application/health/responses"

	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure/factories"
	responseFactory "github.com/JustSteveKing/go-http-response"
)

// Handler is the http.Handler for this service
type Handler struct {
	App *factories.Application
}

// NewHandler will create a new Handler to handle any requests
func NewHandler(app *factories.Application) *Handler {
	return &Handler{
		App: app,
	}
}

// Handle is dispatched when the /ping route is called through the handler
func (handler *Handler) Handle(response http.ResponseWriter, request *http.Request) {
	handler.App.Logger.Info("Service: health-check, Route: /ping")

	responseFactory.Send(
		response,
		http.StatusOK,
		&responses.Response{
			Message: "Service Online",
		},
		handler.App.Config.HTTP.Content,
	)

	return
}
