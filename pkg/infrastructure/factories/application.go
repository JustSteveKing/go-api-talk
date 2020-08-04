package factories

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JustSteveKing/go-api-talk/pkg/infrastructure/providers"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Application is out general purpose Application struct which contains a link through to the main services connected
type Application struct {
	Server *http.Server
	Router *mux.Router
	Logger *zap.Logger
	Config *providers.Config
}

// Run will start the application as required
func (app *Application) Run() {
	err := app.Server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// WaitForShutdown is a graceful way to handle server shutdown events
func (app *Application) WaitForShutdown() {
	// Create a channel to listen for OS signals
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal to our channel
	<-interruptChan

	app.Logger.Info("Received shutdown signal, gracefully terminating")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	app.Server.Shutdown(ctx)
	os.Exit(0)
}
