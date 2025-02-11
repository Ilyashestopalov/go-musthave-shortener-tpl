package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/shortner"
	"github.com/gin-gonic/gin"
)

var (
	serverName string // Server name for listen socket
	baseURL    string // Server name for response
)

// Main function to set up the Gin server
func main() {

	shortener := interfaces.NewMapURLShortener()

	flag.StringVar(&serverName, "a", "localhost:8080", "Server name with port")
	flag.StringVar(&baseURL, "b", "http://localhost:8080", "Base URL for shortened links")

	// Parse the command line flags
	flag.Parse()

	// Overwrite with environment variables if set
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		baseURL = envBaseURL
	}
	if envServerName := os.Getenv("SERVER_NAME"); envServerName != "" {
		serverName = envServerName
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

	router.POST("/", shortner.ShortenURLHandler(shortener, baseURL))
	router.GET("/:shortened", shortner.RedirectURLHandler(shortener))

	router.Run(serverName)
}
