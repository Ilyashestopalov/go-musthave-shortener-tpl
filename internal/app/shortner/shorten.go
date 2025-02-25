package shortner

import (
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/models"
	"github.com/gin-gonic/gin"
)

// Handler for shortening a URL
func ShortenURLHandler(shortener interfaces.URLShortener, baseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		original, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid URL")
			return
		}

		shortened, err := shortener.Shorten(string(original))
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusCreated, baseURL+"/"+shortened)
	}
}

// Handler for shortening a URL via API
func ApiShortenURLHandler(shortener interfaces.URLShortener, baseURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json models.APIShortenURL
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		shortened, err := shortener.Shorten(json.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": baseURL + "/" + shortened})
	}
}
