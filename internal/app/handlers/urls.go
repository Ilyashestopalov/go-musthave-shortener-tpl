package handlers

import (
	"crypto/rand"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/storages"
)

// URLHandler handles URL shortening and retrieval
type URLHandler struct {
	store   storages.DataStore
	baseURL string
}

// NewURLHandler creates a new URLHandler
func NewURLHandler(store storages.DataStore, baseURL string) *URLHandler {
	return &URLHandler{store: store, baseURL: baseURL}
}

// GenerateShortURL generates a random short URL ID
func GenerateShortURL() string {
	const length = 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	_, _ = rand.Read(result)
	for i := range result {
		result[i] = charset[int(result[i])%len(charset)]
	}
	return string(result)
}

/*
// CreateURL creates a new shortened URL
func (h *URLHandler) CreateURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if c.ContentType() == "application/json" {
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
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
*/

/*
// GetURL retrieves the original URL based on the short URL
func (h *URLHandler) GetURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	urlData, exists := h.store.GetURL(shortURL)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.JSON(http.StatusOK, urlData)
}
*/
/*
// DeleteURL deletes a shortened URL based on the short URL
func (h *URLHandler) DeleteURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if err := h.store.DeleteURL(shortURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete URL"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
*/
