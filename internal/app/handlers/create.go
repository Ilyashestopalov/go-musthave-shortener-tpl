package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/generators"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
	"github.com/gin-gonic/gin"
)

// CreateURL creates a new shortened URL
func (h *URLHandler) CreateURL(c *gin.Context) {

	// Handle, TODO move it to sub function
	if c.ContentType() == "application/json" {
		var request struct {
			URL string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		shortURL := generators.SecureRandomString(8)
		urlData := storages.URLData{
			UUID:        fmt.Sprintf("%d", len(h.store.GetAllURLs())+1), // Simple UUID generation based on count
			ShortURL:    shortURL,
			OriginalURL: request.URL,
		}

		if err := h.store.AddURL(urlData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%s/%s", h.baseURL, shortURL)})
	}

	// Handle, TODO move it to sub function
	if c.ContentType() == "text/html; charset=utf-8" {
		// For plain text requests, read the body directly
		var request string
		raw, _ := io.ReadAll(c.Request.Body)
		request = string(raw)
		if request == "" {
			c.String(http.StatusBadRequest, "Invalid URL")
			return
		}

		shortURL := generators.SecureRandomString(8)
		urlData := storages.URLData{
			UUID:        fmt.Sprintf("%d", len(h.store.GetAllURLs())+1), // Simple UUID generation based on count
			ShortURL:    shortURL,
			OriginalURL: request,
		}

		if err := h.store.AddURL(urlData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data"})
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("%s/%s", h.baseURL, shortURL))
		return
	}
}
