package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rsturla/golang-aio/pkg/log"
	"strings"
)

type Config struct {
	Port int `koanf:"port"`
}

func New() *Config {
	return &Config{
		Port: 8080,
	}
}

func (c *Config) Load(filePath string, envPrefix string) error {
	k := koanf.New(".")

	// Load config file.
	if filePath != "" {
		parser := yaml.Parser()
		log.Infof("Loading config file: %s", filePath)
		if err := k.Load(file.Provider(filePath), parser); err != nil {
			log.Errorf("Error loading config file: %v", err)
			return err
		}
	}

	// Load environment variables.
	if err := k.Load(env.Provider(envPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, envPrefix)), "_", ".", -1)
	}), nil); err != nil {
		log.Errorf("Error loading environment variables: %v", err)
	}

	// Unmarshal config file into struct.
	if err := k.Unmarshal("", c); err != nil {
		log.Errorf("Error unmarshalling config file: %v", err)
		return err
	}

	return nil
}
