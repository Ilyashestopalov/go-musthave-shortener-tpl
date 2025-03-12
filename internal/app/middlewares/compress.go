package middlewares

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware compresses the response using gzip for specific content types
/*
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Accept-Encoding") == "" {
			c.Next()
			return
		}

		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		defer gz.Close()

		// Wrap the ResponseWriter
		w := &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		// Process the request
		c.Next()

		// Check if we need to compress the response
		if c.Writer.Status() == http.StatusOK &&
			(c.Writer.Header().Get("Content-Type") == "application/json" ||
				c.Writer.Header().Get("Content-Type") == "text/html") {

			// Close the gzip writer to flush the data
			if err := gz.Close(); err != nil {
				c.Error(err)
				return
			}

			// Set the Content-Encoding header
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Content-Type", c.Writer.Header().Get("Content-Type"))
			c.Writer.WriteHeader(c.Writer.Status())
			io.Copy(c.Writer, &buf)
		}
	}
}
*/

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

		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		defer gz.Close()

		// Wrap the ResponseWriter
		w := &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		// Process the request
		c.Next()

		// Check if we need to compress the response
		if c.Writer.Status() == http.StatusOK &&
			(c.Writer.Header().Get("Content-Type") == "application/json" ||
				c.Writer.Header().Get("Content-Type") == "text/html") {

			// Close the gzip writer to flush the data
			if err := gz.Close(); err != nil {
				c.Error(err)
				return
			}

			// Set the Content-Encoding header
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Content-Type", c.Writer.Header().Get("Content-Type"))
			c.Writer.WriteHeader(c.Writer.Status())
			io.Copy(c.Writer, &buf)
		}
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
