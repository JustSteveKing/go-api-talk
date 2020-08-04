package providers

import "os"

// AppConfig is the Application configuration struct
type AppConfig struct {
	Name    string
	Version string
	Port    string
}

// HTTPConfig is the Application HTTP configuration
type HTTPConfig struct {
	Content string
	Problem string
}

// Config is the Configuration struct
type Config struct {
	App  AppConfig
	HTTP HTTPConfig
}

// ConfigProvider will run and create a new Config struct
func ConfigProvider() *Config {
	return &Config{
		App: AppConfig{
			Name:    env("APP_NAME", "Go App"),
			Version: env("APP_VERSION", "v1.0"),
			Port:    env("APP_PORT", "8080"),
		},
		HTTP: HTTPConfig{
			Content: env("HTTP_CONTENT_TYPE", "application/json"),
			Problem: env("HTTP_PROBLEM", "application/problem+json"),
		},
	}
}

// env is a simple helper function to read an environment variable or return a default value
func env(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
