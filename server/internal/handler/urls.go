package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sota0121/janusly/server/internal/logger"
)

// GenerateTinyURLHandler is a handler for generating tiny URL
func GenerateTinyURLHandler(c *gin.Context) {
	logger.Debugf(c, "generate_tiny_url is called")

	// Get src URL from request.
	srcURL := c.Query("src_url")
	if srcURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "src_url is required",
		})
		return
	}

	// Generate tiny URL.
	// TODO: Implement this.

	logger.Debugf(c, "ping is called")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
