package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sota0121/janusly/server/internal/handler"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.GET("/ping", handler.PingHandler)
	r.GET("/panic", handler.AlwaysPanicHandler)

	return r
}
