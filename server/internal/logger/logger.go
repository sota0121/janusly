package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/sota0121/janusly/server/internal/januslyctx"
)

var Log zerolog.Logger

// SetupLogger sets up logger with zerolog.
func SetupLogger(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

// Infof is a wrapper for zerolog.Info().Msgf.
func Infof(c *gin.Context, format string, v ...interface{}) {
	getLogFromContext(c).Info().Msgf(format, v...)
}

// Debugf is a wrapper for zerolog.Debug().Msgf.
func Debugf(c *gin.Context, format string, v ...interface{}) {
	getLogFromContext(c).Debug().Msgf(format, v...)
}

// Warnf is a wrapper for zerolog.Warn().Msgf.
func Warnf(c *gin.Context, format string, v ...interface{}) {
	getLogFromContext(c).Warn().Msgf(format, v...)
}

// Errorf is a wrapper for zerolog.Error().Msgf.
func Errorf(c *gin.Context, format string, v ...interface{}) {
	getLogFromContext(c).Error().Msgf(format, v...)
}

// getLogFromContext returns logger from gin.Context.
func getLogFromContext(c *gin.Context) *zerolog.Logger {
	l, exists := c.Get(januslyctx.ContextKeyLog.String())
	if exists {
		logger, ok := l.(zerolog.Logger)
		if ok {
			return &logger
		}
	}
	return &Log
}
