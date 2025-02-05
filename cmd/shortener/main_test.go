package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShortenURLHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/", ShortenURLHandler)

	// Test valid request
	longURL := "https://www.example.com"
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(longURL)))
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	if _, exists := response["short_url"]; !exists {
		t.Error("Expected short_url in response")
	}

	// Test invalid request
	invalidLongURL := ""
	req, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(invalidLongURL)))
	req.Header.Set("Content-Type", "text/plain")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, got %d", w.Code)
	}

	var errorResponse map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &errorResponse); err != nil {
		t.Errorf("Failed to parse error response body: %v", err)
	}

	if _, exists := errorResponse["error"]; !exists {
		t.Error("Expected error in response")
	}
}

func TestRedirectHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/", ShortenURLHandler)
	router.GET("/:short_url", RedirectHandler)

	// First create a short URL
	longURL := "https://www.example.com"
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(longURL)))
	req.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}
	shortURL := response["short_url"][len(baseURL):] // Extract the short URL identifier

	// Now test redirect
	req, _ = http.NewRequest(http.MethodGet, "/"+shortURL, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code 301, got %d", w.Code)
	}
}
