package log

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	logger       *logrus.Logger
	logLevelLock = sync.Mutex{}
)

type LogLevel int
type FormatterType int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

const (
	JSONFormatter FormatterType = iota
	TextFormatter
)

// SetLogLevel sets the log level of the logger.
func SetLogLevel(level LogLevel) {
	logLevelLock.Lock()
	defer logLevelLock.Unlock()

	logger.SetLevel(logrus.Level(level))
}

// SetFormatter sets the log formatter of the logger.
func SetFormatter(formatter FormatterType) {
	switch formatter {
	case JSONFormatter:
		logger.SetFormatter(&logrus.JSONFormatter{})
	case TextFormatter:
		logger.SetFormatter(&logrus.TextFormatter{})
	}
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
