package middlewares

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware checks for Content-Encoding and compresses responses
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the Content-Encoding header is set to gzip
		if c.Request.Header.Get("Content-Encoding") == "gzip" {
			// Create a new gzip reader to decompress the request body
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid gzip encoding")
				c.Abort()
				return
			}
			defer reader.Close()

			// Read the decompressed body
			decompressedBody, err := io.ReadAll(reader)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error reading decompressed body")
				c.Abort()
				return
			}

			// Replace the original request body with the decompressed data
			c.Request.Body = io.NopCloser(bytes.NewBuffer(decompressedBody))
			c.Request.ContentLength = int64(len(decompressedBody))
		}

		// Create a buffer to hold the compressed response
		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		defer gz.Close()

		// Wrap the ResponseWriter to capture the response
		w := &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		// Process the request
		c.Next()

		// Check if the response status is OK and if the request accepts gzip encoding
		if c.Writer.Status() == http.StatusOK && c.Request.Header.Get("Accept-Encoding") == "gzip" {
			// Check if the Content-Type is application/json or text/html
			contentType := c.Writer.Header().Get("Content-Type")
			if contentType == "application/json" || contentType == "text/html" {
				// Close the gzip writer to flush the data
				if err := gz.Close(); err != nil {
					c.Error(err)
					return
				}

				// Set the Content-Encoding header to gzip
				c.Writer.Header().Set("Content-Encoding", "gzip")
				c.Writer.Header().Set("Content-Type", contentType) // Maintain original Content-Type
				c.Writer.WriteHeader(c.Writer.Status())
				io.Copy(c.Writer, &buf)
			}
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
