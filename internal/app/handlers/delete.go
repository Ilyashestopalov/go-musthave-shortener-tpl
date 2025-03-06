package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteURL deletes a shortened URL based on the short URL
func (h *URLHandler) DeleteURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if err := h.store.DeleteURL(shortURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete URL"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
