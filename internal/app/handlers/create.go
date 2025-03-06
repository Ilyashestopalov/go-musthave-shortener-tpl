package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
	"github.com/gin-gonic/gin"
)

// CreateURL creates a new shortened URL
func (h *URLHandler) CreateURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	// Handle
	if c.ContentType() == "application/json" {
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Handle
	if c.ContentType() == "text/plain" {
		// For plain text requests, read the body directly
		request.URL = c.PostForm("url")
		if request.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
			return
		}
	}

	shortened := GenerateShortURL()
	urlData := storages.URLData{
		UUID:        fmt.Sprintf("%d", len(h.store.GetAllURLs())+1), // Simple UUID generation based on count
		ShortURL:    shortened,
		OriginalURL: request.URL,
	}

	if err := h.store.AddURL(urlData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": fmt.Sprintf("%s/%s", h.baseURL, shortened)})
}
