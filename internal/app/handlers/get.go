package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetURL retrieves the original URL based on the short URL
func (h *URLHandler) GetURL(c *gin.Context) {
	if c.ContentType() == "text/html; charset=utf-8" {
		shortURL := c.Param("short_url")
		urlData, exists := h.store.GetURL(shortURL)
		if !exists {
			c.String(http.StatusNotFound, "URL Not Found")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, urlData.OriginalURL)
	}
}
