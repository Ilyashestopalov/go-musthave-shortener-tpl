package main

import (
	"log"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/server"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"go.uber.org/zap"
)

// Main function to set up the Gin server
func main() {

	// Initialize the logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
	defer logger.Sync() // Flush any buffered log entries

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	err = server.Run(logger, cfg)
	if err != nil {
		log.Fatal("Server failed to start", err)
	}

}
