package januslyctx

import (
	"github.com/gin-gonic/gin"
)

// ContextKey is a type for context key.
type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

// ContextKeyLog is a context key for logger.
const ContextKeyLog ContextKey = "log"

// ContextKeyTraceID is a context key for trace id.
const ContextKeyTraceID ContextKey = "trace-id"

func SetTraceID(c *gin.Context, traceID string) {
	c.Set(ContextKeyTraceID.String(), traceID)
}

func GetTraceID(c *gin.Context) string {
	traceID, exists := c.Get(ContextKeyTraceID.String())
	if exists {
		return traceID.(string)
	}
	return ""
}

func SetLogger(c *gin.Context, logger interface{}) {
	c.Set(ContextKeyLog.String(), logger)
}
