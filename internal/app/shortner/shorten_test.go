package shortner

import (
	"bytes"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	serverName string // Server name for listen socket
	baseURL    string // Server name for response
)

func setupRouter(shortener interfaces.URLShortener, baseURL string) *gin.Engine {
	r := gin.Default()
	r.POST("/", ShortenURLHandler(shortener, baseURL))
	r.GET("/:shortened", RedirectURLHandler(shortener))
	return r
}

func TestURLShortener(t *testing.T) {

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

	shortener := interfaces.NewMapURLShortener()
	router := setupRouter(shortener, baseURL)

	t.Run("Shorten URL", func(t *testing.T) {
		reqBody := `{"url": "https://www.example.com"}`
		req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer([]byte(reqBody)))
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Body.String(), "shortened_url")
	})
}
