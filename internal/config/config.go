package config

import (
	"flag"
	"os"
)

var (
	serverName string
	baseURL    string
)

// Config holds the configuration for the server
type Config struct {
	BaseURL    string
	ServerName string
}

func LoadConfig() (*Config, error) {
	flag.StringVar(&serverName, "a", "localhost:8080", "Server name with port")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "Base URL for shortened links")

	// Parse the command line flags
	flag.Parse()

	// Overwrite with environment variables if set
	if baseURL == "" {
		baseURL = os.Getenv("BASE_URL")
	}
	if serverName == "" {
		serverName = os.Getenv("SERVER_NAME")
	}

	return &Config{
		BaseURL:    baseURL,
		ServerName: serverName,
	}, nil
}
