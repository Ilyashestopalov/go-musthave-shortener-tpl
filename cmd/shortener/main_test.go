package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/", ShortenURLHandler)
	router.GET("/:short_url", RedirectHandler)
	return router
}

func TestShortenURLHandler(t *testing.T) {
	router := setupRouter()

	// Test valid request
	longURL := "https://www.example.com"
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(longURL)))
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Test invalid request
	invalidLongURL := ""
	req, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(invalidLongURL)))
	req.Header.Set("Content-Type", "text/plain")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
