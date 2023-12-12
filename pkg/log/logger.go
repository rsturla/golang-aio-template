package log

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	logger       *logrus.Logger
	logLevelLock = sync.Mutex{}
)

// SetLogLevel sets the log level of the logger.
func SetLogLevel(level logrus.Level) {
	logLevelLock.Lock()
	defer logLevelLock.Unlock()

	logger.SetLevel(level)
}

// SetFormatter sets the log formatter of the logger.
func SetFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}

// Logger returns the current logger instance.
func Logger() *logrus.Logger {
	return logger
}

// SetLogger sets a custom logger instance.
func SetLogger(l *logrus.Logger) {
	logger = l
}

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}
