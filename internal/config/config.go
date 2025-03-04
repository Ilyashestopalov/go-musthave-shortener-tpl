package config

import (
	"flag"
	"os"
)

var (
	serverName      string
	baseURL         string
	fileStoragePath string
)

// Config holds the configuration for the server
type Config struct {
	BaseURL         string
	ServerName      string
	FileStoragePath string
}

func LoadConfig() (*Config, error) {
	flag.StringVar(&serverName, "a", "localhost:8080", "Server name with port")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "Base URL for shortened links")
	flag.StringVar(&fileStoragePath, "f", "/tmp/data.json", "Path for storage file")

	// Parse the command line flags
	flag.Parse()

	// Overwrite with environment variables if set
	if v := os.Getenv("BASE_URL"); v != "" {
		baseURL = v
	}
	if v := os.Getenv("SERVER_NAME"); v != "" {
		serverName = v
	}
	if v := os.Getenv("FILE_STORAGE_PATH"); v != "" {
		fileStoragePath = v
	}

	/*
		if os.Getenv("BASE_URL") != "" {
			baseURL = os.Getenv("BASE_URL")
		}

		if os.Getenv("SERVER_NAME") != "" {
			serverName = os.Getenv("SERVER_NAME")
		}

		if os.Getenv("FILE_STORAGE_PATH") != "" {
			fileStoragePath = os.Getenv("FILE_STORAGE_PATH")
		}
	*/

	return &Config{
		BaseURL:         baseURL,
		ServerName:      serverName,
		FileStoragePath: fileStoragePath,
	}, nil
}
