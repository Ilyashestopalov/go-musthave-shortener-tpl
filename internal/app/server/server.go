package server

import (
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/handlers"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/middlewares"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/services"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger, cfg *config.Config) error {
	gin.SetMode(gin.ReleaseMode)

	defer logger.Sync()

	router := gin.New()

	sugar := logger.Sugar()
	router.Use(middlewares.LoggerMiddleware(sugar))
	router.Use(middlewares.GzipMiddleware())

	store := storages.NewFileStore(cfg.FileStoragePath)
	service := services.NewURLService(store)
	handler := handlers.NewHandler(cfg, service)

	router.GET("/:url", handler.RedirectURL)
	router.POST("/", handler.URLCreator)
	router.POST("/api/shorten", handler.URLCreatorJSON)
	// fmt.Printf("%+v\n", cfg)
	return router.Run(cfg.ServerName)
}
