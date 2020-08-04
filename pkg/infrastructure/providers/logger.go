package providers

import "go.uber.org/zap"

// LoggerProvider creates and returns a Logger for our application
func LoggerProvider() *zap.Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	return logger
}
