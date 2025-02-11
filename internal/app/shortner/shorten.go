package shortner

import (
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
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
