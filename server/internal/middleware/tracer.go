package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sota0121/janusly/server/internal/januslyctx"
)

// TracerMiddleware is a middleware for tracing.
// This middleware sets trace-id for tracking request.
// Caution: This middleware must be set before LoggerMiddleware. Because LoggerMiddleware uses trace-id.
func TracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.NewString()
		}
		januslyctx.SetTraceID(c, traceID)

		c.Next()
	}
}
