package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/sota0121/janusly/server/internal/logger"
	"github.com/sota0121/janusly/server/internal/middleware"
	"github.com/sota0121/janusly/server/internal/router"
)

func main() {
	// Initialize logger
	// Note: You must initialize logger before setup middleware.
	logger.SetupLogger(zerolog.DebugLevel)

	r := gin.New()

	// Setup middleware and router
	// Note: You must set middleware before router.
	r = middleware.SetupMiddleware(r)
	r = router.SetupRouter(r)

	r.Run() // listen and serve on
}
