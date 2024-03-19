package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sota0121/janusly/server/internal/handler"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.GET("/ping", handler.PingHandler)
	r.GET("/panic", handler.AlwaysPanicHandler)

	// URL Management
	urls := r.Group("/urls")
	{
		urls.GET("/generate_tiny_url", handler.GenerateTinyURLHandler)
	}

	return r
}
