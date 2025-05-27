package configs

import (
	"flag"
	"os"
)

// Config holds the configuration parameters for the application
type Config struct {
	FileStoragePath string
	PostgresURL     string
	ServerName      string
	BaseURL         string
}

// getEnv read env variable or return fallback (default)
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// LoadConfig loads configuration from environment variables or command-line arguments!
func LoadConfig() *Config {
	filePath := flag.String("f", getEnv("FILE_STORAGE_PATH", "/tmp/storage.data"), "File storage path")
	postgresURL := flag.String("d", getEnv("POSTGRES_STORAGE_URL", ""), "PostgreSQL connection URL")
	serverName := flag.String("a", getEnv("SERVER_NAME", "localhost:8080"), "Server address")
	baseURL := flag.String("b", getEnv("BASE_URL", "http://localhost:8080"), "Base URL for redirect")

	flag.Parse()

	return &Config{
		FileStoragePath: *filePath,
		PostgresURL:     *postgresURL,
		ServerName:      *serverName,
		BaseURL:         *baseURL,
	}
}
