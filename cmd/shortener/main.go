package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var (
	urlStore   = make(map[string]string) // Store for mapping short URLs to long URLs
	serverName string                    // Server name for listen socket
	baseURL    string                    // Server name for response
)

// ShortenURLHandler handles the URL shortening requests
func ShortenURLHandler(c *gin.Context) {
	longURL, err := c.GetRawData()
	if err != nil || len(longURL) == 0 {
		c.String(http.StatusBadRequest, "Invalid URL")
		return
	}

	shortURL := generateShortURL()
	urlStore[shortURL] = string(longURL)
	c.String(http.StatusCreated, baseURL+"/"+shortURL)
}

// RedirectHandler redirects short URL to the original long URL
func RedirectHandler(c *gin.Context) {
	shortURL := c.Param("short_url")

	longURL, exists := urlStore[shortURL]

	if !exists {
		c.String(http.StatusNotFound, "URL not found")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, longURL)
}

// TODO: move it internal/app/helpers/generators.go
func generateShortURL() string {
	var charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var output strings.Builder
	length := 8
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func main() {
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
	//router.Use(gin.Recovery())

	router.POST("/", ShortenURLHandler)
	router.GET("/:short_url", RedirectHandler)

	router.Run(serverName)
}
