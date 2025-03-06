package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggingMiddleware logs requests using Zap logger
func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Request received",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)

		c.Next()

		logger.Info("Response sent",
			zap.String("status", http.StatusText(c.Writer.Status())),
			zap.String("path", c.Request.URL.Path),
		)
	}
}
