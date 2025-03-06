package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetURL retrieves the original URL based on the short URL
func (h *URLHandler) GetURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	urlData, exists := h.store.GetURL(shortURL)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found Bleat"})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, urlData.OriginalURL)
}
