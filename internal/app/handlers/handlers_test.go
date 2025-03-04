package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockService struct{}

func (m *mockService) ShortenURL(url string) (string, error) {
	if url == "http://example.com" {
		return "xxxxxx", nil
	}
	return "", errors.New("invalid URL format")
}

func (m *mockService) GetOriginalURL(input string) (string, bool) {
	if input == "xxxxxx" {
		return "http://example.com", true
	}
	return "", false
}

func TestURLCreator(t *testing.T) {
	cfg := &config.Config{BaseURL: "http://localhost:8080"}
	handler := NewHandler(cfg, &mockService{})

	router := gin.New()
	router.POST("/", handler.URLCreator)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("http://example.com"))
	r.Header.Set("Content-Type", "text/plain")

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotNil(t, res)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestRedirectURL(t *testing.T) {
	cfg := &config.Config{BaseURL: "http://localhost:8080"}
	handler := NewHandler(cfg, &mockService{})

	router := gin.New()
	router.GET("/:url", handler.RedirectURL)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/xxxxxx", nil)

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotNil(t, res)
	assert.Equal(t, http.StatusTemporaryRedirect, res.StatusCode)
	assert.Equal(t, "http://example.com", res.Header.Get("Location"))
}

func TestURLCreatorJSON(t *testing.T) {
	cfg := &config.Config{BaseURL: "http://localhost:8080"}
	handler := NewHandler(cfg, &mockService{})

	router := gin.New()
	router.POST("/api/shorten", handler.URLCreatorJSON)

	w := httptest.NewRecorder()
	jsonBody := `{"url": "http://example.com"}`
	r := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(jsonBody))
	r.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.NotNil(t, res)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	expectedResponse := `{"result":"http://localhost:8080/xxxxxx"}`
	bodyBytes, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	assert.JSONEq(t, expectedResponse, string(bodyBytes))
}
