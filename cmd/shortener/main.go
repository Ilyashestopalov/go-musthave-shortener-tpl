package main

import (
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/configs"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/server"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"

	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := configs.LoadConfig()

	// Initialize the logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger")
	}
	defer logger.Sync()
	/*
		if cfg.PostgresURL != "" {
			store, err := storages.NewPostgreSQLStore(cfg.PostgresURL)
			if err != nil {
				logger.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
			}
			server.StartServer(store, logger, cfg)
	/*/
	if cfg.FileStoragePath != "" {
		store := storages.NewFileStore(cfg.FileStoragePath)
		if err := store.LoadData(); err != nil {
			logger.Fatal("Failed to load data", zap.Error(err))
		}
		server.StartServer(store, logger, cfg)
	} else {
		store := storages.NewInMemoryStore()
		server.StartServer(store, logger, cfg)
	}

	store := storages.NewInMemoryStore()
	server.StartServer(store, logger, cfg)
}
