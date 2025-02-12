package shortner

import (
	"net/http"

	"github.com/Ilyashestopalov/go-musthave-shortener-tpl/internal/app/interfaces"
	"github.com/gin-gonic/gin"
)

// RedirectHandler redirects short URL to the original long URL
func RedirectURLHandler(shortener interfaces.URLShortener) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortened := c.Param("shortened")
		original, err := shortener.Retrieve(shortened)
		if err != nil {
			c.String(http.StatusNotFound, "URL Not Found")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, original)
	}
}
