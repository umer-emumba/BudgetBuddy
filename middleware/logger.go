package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinZapLogger is a middleware function that logs each request using Zap logger.
func GinZapLogger(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Request details
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// Log using Zap logger
		logger.Info("Request",
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
		)

		// Check if an error occurred during processing the request
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error("Request failed",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.Error(err),
				)
			}
		}
	}
}
