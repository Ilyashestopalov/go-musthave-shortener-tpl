package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetURL retrieves the original URL based on the short URL
func (h *URLHandler) GetURL(c *gin.Context) {
	if c.ContentType() != "application/json" {
		shortURL := c.Param("short_url")
		if shortURL == "" {
			c.String(http.StatusOK, "")
			return
		} else {
			urlData, exists := h.store.GetURL(shortURL)
			if !exists {
				c.String(http.StatusNotFound, "URL Not Found")
				return
			}
			c.Redirect(http.StatusTemporaryRedirect, urlData.OriginalURL)
		}
	}
}
