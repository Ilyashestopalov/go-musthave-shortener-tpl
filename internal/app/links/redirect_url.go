package links

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedirectHandler redirects short URL to the original long URL
func RedirectHandler(urlStore map[string]string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		shortURL := c.Param("short_url")

		longURL, exists := urlStore[shortURL]

		if !exists {
			c.String(http.StatusNotFound, "URL not found")
		} else {
			c.Redirect(http.StatusTemporaryRedirect, longURL)
		}
	}
	return gin.HandlerFunc(fn)
}
