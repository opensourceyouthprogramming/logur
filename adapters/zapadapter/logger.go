// Package zapadapter provides a logur compatible adapter for Uber's Zap.
package zapadapter

import (
	"github.com/goph/logur"
	"github.com/goph/logur/internal/keyvals"
	"go.uber.org/zap"
)

type adapter struct {
	logger *zap.SugaredLogger
}

// New returns a new logur compatible logger with zap as the logging library.
// If nil is passed as logger, the global sugared logger instance is used as fallback.
func New(logger *zap.SugaredLogger) logur.Logger {
	if logger == nil {
		logger = zap.S()
	}

	return &adapter{logger}
}

func (a *adapter) Trace(msg string, fields map[string]interface{}) {
	// Fall back to Debug
	a.Debug(msg, fields)
}

func (a *adapter) Debug(msg string, fields map[string]interface{}) {
	a.logger.Debugw(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Info(msg string, fields map[string]interface{}) {
	a.logger.Infow(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Warn(msg string, fields map[string]interface{}) {
	a.logger.Warnw(msg, keyvals.FromMap(fields)...)
}

func (a *adapter) Error(msg string, fields map[string]interface{}) {
	a.logger.Errorw(msg, keyvals.FromMap(fields)...)
}
