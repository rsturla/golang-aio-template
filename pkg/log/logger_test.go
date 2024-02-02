package log

import (
	"github.com/sirupsen/logrus"
	"testing"
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

			if (err != nil) != tt.err {
				t.Errorf("SetLogLevel() error = %v, wantErr %v", err, tt.err)
			}

			if logger.GetLevel() != tt.want {
				t.Errorf("SetLogLevel() got = %v, want %v", logger.GetLevel(), tt.want)
			}
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
				if !ok || formatter == nil {
					t.Error("SetFormatter() failed to set JSON formatter")
				}
			case TextFormatter:
				formatter, ok := logger.Formatter.(*logrus.TextFormatter)
				if !ok || formatter == nil {
					t.Error("SetFormatter() failed to set Text formatter")
				}
			}
		})
	}
}

func TestLogger(t *testing.T) {
	logger = logrus.New()
	result := Logger()

	if result != logger {
		t.Error("Logger() returned incorrect logger instance")
	}
}

func TestSetLogger(t *testing.T) {
	customLogger := logrus.New()
	SetLogger(customLogger)

	if logger != customLogger {
		t.Error("SetLogger() failed to set custom logger instance")
	}
}
