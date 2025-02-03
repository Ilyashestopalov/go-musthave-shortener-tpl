package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
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

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "short_url")

	// Test invalid request
	invalidLongURL := ""
	req, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(invalidLongURL)))
	req.Header.Set("Content-Type", "text/plain")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestRedirectHandler(t *testing.T) {
	router := setupRouter()

	// First create a short URL
	longURL := "https://www.example.com"
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(longURL)))
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	shortURL := response["short_url"][len(baseURL):] // Extract the short URL identifier

	// Now test redirect
	req, _ = http.NewRequest(http.MethodGet, "/"+shortURL, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}
