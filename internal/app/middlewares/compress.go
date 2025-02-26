package middlewares

import (
	"bytes"
	"compress/gzip"
	"io"

	"github.com/gin-gonic/gin"
)

// gzipResponseWriter is a wrapper for gin's ResponseWriter to write compressed data
type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

// GzipMiddleware compresses the response using gzip
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Accept-Encoding") == "" &&
			(c.Request.Header.Get("Content-Type") != "application/json" || c.Request.Header.Get("Content-Type") != "text/html") {
			c.Next()
			return
		}
		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)

		// Wrap the ResponseWriter
		w := &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		// Process request
		c.Next()

		// Close the gzip writer to flush the data
		if err := gz.Close(); err != nil {
			c.Error(err)
			return
		}

		// Set the Content-Encoding header
		c.Writer.Header().Set("Content-Encoding", "gzip")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(c.Writer.Status())
		io.Copy(c.Writer, &buf)
	}
}
