package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sota0121/janusly/server/internal/januslyctx"
	"github.com/sota0121/janusly/server/internal/logger"
)

// LoggerMiddleware is a middleware for logging.
// This middleware sets logger object to gin.Context with setting necessary fields.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prepare information for logging.
		traceID := januslyctx.GetTraceID(c)

		// Set up logger.
		requestLogger := logger.Log.With().
			Str("url", c.Request.URL.String()).
			Str("method", c.Request.Method).
			Str("trace-id", traceID).
			Logger()

		// Set logger to gin.Context.
		januslyctx.SetLogger(c, requestLogger)

		c.Next()
	}
}
