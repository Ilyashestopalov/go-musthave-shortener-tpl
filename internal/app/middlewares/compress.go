package middlewares

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware compresses the response using gzip for specific content types
func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the request has Content-Encoding set to gzip
		if c.Request.Header.Get("Content-Encoding") == "gzip" {
			// Decompress the body for further processing
			reader, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.String(http.StatusBadRequest, "Invalid gzip encoding")
				c.Abort()
				return
			}
			defer reader.Close()

			// Replace the request body with the decompressed body
			decompressedBody, err := io.ReadAll(reader)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error reading decompressed body")
				c.Abort()
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(decompressedBody))
			c.Request.ContentLength = int64(len(decompressedBody))
		}

		// Create a buffer to hold the compressed response
		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		defer gz.Close()

		// Wrap the ResponseWriter
		w := &gzipResponseWriter{ResponseWriter: c.Writer, Writer: gz}
		c.Writer = w

		// Process the request
		c.Next()

		// Check if we need to compress the response
		if c.Writer.Status() == http.StatusOK && (c.Writer.Header().Get("Accept-Encoding") == "gzip" || c.Writer.Header().Get("Content-Encoding") == "gzip") {
			// Close the gzip writer to flush the data
			if err := gz.Close(); err != nil {
				c.Error(err)
				return
			}

			// Set the Content-Encoding header
			c.Writer.Header().Set("Content-Encoding", "gzip")
			c.Writer.Header().Set("Accept-Encoding", "gzip")
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
