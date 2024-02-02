package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  logrus.Level
		err   bool
	}{
		{"ValidLevel", "debug", logrus.DebugLevel, false},
		{"InvalidLevel", "invalid", logrus.DebugLevel, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger = logrus.New()
			err := SetLogLevel(tt.level)

			assert.Equal(t, tt.err, err != nil, "SetLogLevel() error mismatch")
			assert.Equal(t, tt.want, logger.GetLevel(), "SetLogLevel() log level mismatch")
		})
	}
}

func TestSetFormatter(t *testing.T) {
	tests := []struct {
		name      string
		formatter FormatterType
		want      logrus.Formatter
	}{
		{"JSONFormatter", JSONFormatter, &logrus.JSONFormatter{}},
		{"TextFormatter", TextFormatter, &logrus.TextFormatter{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger = logrus.New()
			SetFormatter(tt.formatter)

			switch tt.formatter {
			case JSONFormatter:
				formatter, ok := logger.Formatter.(*logrus.JSONFormatter)
				assert.True(t, ok, "SetFormatter() failed to set JSON formatter")
				assert.NotNil(t, formatter, "SetFormatter() failed to set JSON formatter")
			case TextFormatter:
				formatter, ok := logger.Formatter.(*logrus.TextFormatter)
				assert.True(t, ok, "SetFormatter() failed to set Text formatter")
				assert.NotNil(t, formatter, "SetFormatter() failed to set Text formatter")
			}
		})
	}
}

func TestLogger(t *testing.T) {
	logger = logrus.New()
	result := Logger()

	assert.Equal(t, logger, result, "Logger() returned incorrect logger instance")
}

func TestSetLogger(t *testing.T) {
	customLogger := logrus.New()
	SetLogger(customLogger)

	assert.Equal(t, customLogger, logger, "SetLogger() failed to set custom logger instance")
}
