package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(store storages.DataStore) *gin.Engine {
	r := gin.Default()
	urlHandler := NewURLHandler(store, "http://localhost:8080")

	r.POST("/", urlHandler.CreateURL)
	r.GET("/:short_url", urlHandler.GetURL)
	r.DELETE("/:short_url", urlHandler.DeleteURL)

	return r
}

func TestURLShortener(t *testing.T) {
	store := storages.NewInMemoryStore()
	router := setupRouter(store)
	/*
		// Test POST with JSON
		t.Run("Create URL with JSON", func(t *testing.T) {
			urlToShorten := "http://example.com"
			jsonValue, _ := json.Marshal(map[string]string{"url": urlToShorten})
			req, _ := http.NewRequest("POST", "/urls", bytes.NewBuffer(jsonValue))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusCreated, resp.Code)

			var response map[string]string
			json.Unmarshal(resp.Body.Bytes(), &response)

			shortURL := response["result"]
			assert.NotEmpty(t, shortURL)

			// Test GET
			req, _ = http.NewRequest("GET", "/urls/"+shortURL, nil)
			resp = httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusOK, resp.Code)

			var retrievedResponse URLData
			json.Unmarshal(resp.Body.Bytes(), &retrievedResponse)

			assert.Equal(t, shortURL, retrievedResponse.ShortURL)
			assert.Equal(t, urlToShorten, retrievedResponse.OriginalURL)

			// Test DELETE
			req, _ = http.NewRequest("DELETE", "/urls/"+shortURL, nil)
			resp = httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNoContent, resp.Code)

			// Test GET after delete
			req, _ = http.NewRequest("GET", "/urls/"+shortURL, nil)
			resp = httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, http.StatusNotFound, resp.Code)
		})
	*/
	// Test POST with plain text
	t.Run("Create URL with plain text", func(t *testing.T) {
		urlToShorten := "http://example2.com"
		// Test POST
		req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(urlToShorten))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)

		var response string

		shortURL := response
		req, _ = http.NewRequest("GET", "/"+shortURL, nil)
		resp = httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)

		var retrievedResponse storages.URLData
		json.Unmarshal(resp.Body.Bytes(), &retrievedResponse)

		assert.Equal(t, shortURL, retrievedResponse.ShortURL)

	})
}
