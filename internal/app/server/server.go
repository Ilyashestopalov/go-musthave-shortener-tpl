package server

import (
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/configs"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/handlers"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/middlewares"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// StartServer initializes the server and routes
func StartServer(store storages.DataStore, logger *zap.Logger, cfg *configs.Config) {
	router := gin.Default()
	router.Use(middlewares.LoggingMiddleware(logger))
	router.Use(middlewares.GzipMiddleware())

	urlHandler := handlers.NewURLHandler(store, cfg.BaseURL)

	router.POST("/", urlHandler.CreateURL)
	router.GET("/:short_url", urlHandler.GetURL)
	router.DELETE("/:short_url", urlHandler.DeleteURL)

	// Start the server
	if err := router.Run(cfg.ServerName); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
