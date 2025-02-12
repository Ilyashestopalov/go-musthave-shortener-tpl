package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/shortner"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"github.com/gin-gonic/gin"
)

// Main function to set up the Gin server
func main() {
	shortener := interfaces.NewMapURLShortener()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s ;; %s ;; %s %s ;; %s ;; %d ;; %s ;; %s ;; %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.POST("/", shortner.ShortenURLHandler(shortener, cfg.BaseURL))
	router.GET("/:shortened", shortner.RedirectURLHandler(shortener))

	router.Run(cfg.ServerName)
}
