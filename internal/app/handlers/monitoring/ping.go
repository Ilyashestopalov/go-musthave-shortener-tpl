package monitoring

import (
	"github.com/gin-gonic/gin"
)

// GetURL retrieves the original URL based on the short URL
func GetPing(c *gin.Context) {

	c.Header("Content-Type", "text/plain")

}
