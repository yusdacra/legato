package logger

import (
	"os"
	"runtime/debug"

	"github.com/harmony-development/legato/server/config"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type ILogger interface {
	CheckException(err error)
	Exception(err error)
	Fatal(err error)
}

// Logger is the Harmony logger
type Logger struct {
	Config *config.Config
}

// New creates a Logger with a given configuration
func New(cfg *config.Config) *Logger {
	return &Logger{
		Config: cfg,
	}
}

// CheckException logs an exception if it's defined
func (l Logger) CheckException(err error) {
	if err == nil {
		return
	}
	l.Exception(err)
}

// Exception logs an exception
func (l Logger) Exception(err error) {
	if l.Config.Sentry.Enabled {
		sentry.CaptureException(err)
	}
	if l.Config.Server.LogErrors {
		logrus.Warn(err, string(debug.Stack()))
	}
}

// Fatal logs an exception and then aborts
func (l Logger) Fatal(err error) {
	l.CheckException(err)
	os.Exit(-1)
}
