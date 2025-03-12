package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware compresses the response using gzip for specific content types
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Gzip content"})
				return
			}
			defer reader.Close()
			c.Request.Body = io.NopCloser(reader)
		}

		if !strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		gzWriter := gzip.NewWriter(c.Writer)
		defer gzWriter.Close()

		c.Writer = &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gzWriter}
		c.Header("Content-Encoding", "gzip")

		c.Next()
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
