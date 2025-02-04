package links

import (
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/helpers"
	"github.com/gin-gonic/gin"
)

// ShortenURLHandler handles the URL shortening requests
func ShortenURLHandler(baseURL string, urlStore map[string]string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		longURL, err := c.GetRawData()
		if err != nil || len(longURL) == 0 {
			c.String(http.StatusBadRequest, "Invalid URL")
		} else {
			shortURL := helpers.GenerateId()
			urlStore[shortURL] = string(longURL)
			c.String(http.StatusCreated, baseURL+"/"+shortURL)
		}
	}
	return gin.HandlerFunc(fn)

}
