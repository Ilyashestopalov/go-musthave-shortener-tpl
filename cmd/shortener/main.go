package main

import (
	"log"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/middlewares"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/routes"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"github.com/gin-gonic/gin"
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

	shortener := interfaces.NewMapURLShortener()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middlewares.Logger(logger))

	// Register routes
	routes.RegisterRoutes(router, shortener, cfg.BaseURL)

	// Run server
	router.Run(cfg.ServerName)
}
