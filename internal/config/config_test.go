package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		envVars     map[string]string
		expected    Config
	}{
		{
			name:        "DefaultValues",
			fileContent: "",
			envVars:     nil,
			expected: Config{
				Port: 8080,
			},
		},
		{
			name:        "LoadConfigFromFile",
			fileContent: "port: 9090",
			envVars:     nil,
			expected: Config{
				Port: 9090,
			},
		},
		{
			name:        "LoadConfigFromEnv",
			fileContent: "",
			envVars: map[string]string{
				"MYAPP_PORT": "7070",
			},
			expected: Config{
				Port: 7070,
			},
		},
		{
			name:        "LoadConfigFromFileAndEnv",
			fileContent: "port: 6060",
			envVars: map[string]string{
				"MYAPP_PORT": "5050",
			},
			expected: Config{
				Port: 5050,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare a temporary YAML config file
			filePath := "testconfig.yaml"
			defer os.Remove(filePath)

			err := os.WriteFile(filePath, []byte(tt.fileContent), 0644)
			assert.NoError(t, err, "Error writing test config file")

			// Set environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			// Load configuration
			cfg := New()
			err = cfg.Load(filePath, "MYAPP_")
			assert.NoError(t, err, "Error loading config")

			// Assert values
			assert.Equal(t, tt.expected.Port, cfg.Port, "Port value mismatch")
		})
	}
}
