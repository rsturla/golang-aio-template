package log

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	logger       *logrus.Logger
	logLevelLock = sync.Mutex{}
)

func SetLogLevel(level logrus.Level) {
	defer logLevelLock.Unlock()
	logLevelLock.Lock()

	logger.Level = level
}

func SetFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

func Logger() *logrus.Logger {
	return logger
}

func SetLogger(l *logrus.Logger) {
	logger = l
}

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}
