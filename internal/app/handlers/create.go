package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/generators"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
	"github.com/gin-gonic/gin"
)

// CreateURL creates a new shortened URL
func (h *URLHandler) CreateURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	// Handle, TODO move it to sub function
	if c.ContentType() == "application/json" {
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
		c.JSON(http.StatusCreated, gin.H{"result": fmt.Sprintf("%s/%s", h.baseURL, shortURL)})
	} else {
		// For plain text requests, read the body directly
		r, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid URL")
			return
		}

		shortURL := generators.SecureRandomString(8)
		urlData := storages.URLData{
			UUID:        fmt.Sprintf("%d", len(h.store.GetAllURLs())+1), // Simple UUID generation based on count
			ShortURL:    shortURL,
			OriginalURL: fmt.Sprintf("%s", r),
		}

		if err := h.store.AddURL(urlData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save data"})
			return
		}
		fmt.Printf("ebolaaaaaaaa!!!!!!")
		c.String(http.StatusCreated, fmt.Sprintf("%s/%s", h.baseURL, shortURL))
		return
	}
}
