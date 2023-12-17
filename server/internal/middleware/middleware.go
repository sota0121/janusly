package middleware

import (
	"github.com/gin-gonic/gin"
)

// SetupMiddleware sets up middleware for gin.Engine.
func SetupMiddleware(r *gin.Engine) *gin.Engine {
	// Set middleware you want to apply first.
	r.Use(TracerMiddleware())

	// Set the other middleware.
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// These middlewares must be set after TracerMiddleware.
	// Because these middlewares use trace-id.
	r.Use(LoggerMiddleware())

	return r
}
