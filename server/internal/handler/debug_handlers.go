package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sota0121/janusly/server/internal/logger"
)

// PingHandler is a handler for ping.
func PingHandler(c *gin.Context) {
	logger.Debugf(c, "ping is called")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// AlwaysPanicHandler is a handler for panic.
func AlwaysPanicHandler(c *gin.Context) {
	logger.Debugf(c, "panic is called")
	panic("always panic")
}
