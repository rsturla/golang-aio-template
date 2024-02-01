package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rsturla/golang-aio/pkg/log"
)

type Config struct {
	Port int `koanf:"port"`
}

func New() *Config {
	return &Config{
		Port: 8080,
	}
}

func (c *Config) Load(filePath string) error {
	k := koanf.New(".")
	parser := yaml.Parser()
	log.Infof("Loading config file: %s", filePath)

	if err := k.Load(file.Provider(filePath), parser); err != nil {
		log.Errorf("Error loading config file: %v", err)
		return err
	}

	if err := k.Unmarshal("", c); err != nil {
		log.Errorf("Error unmarshalling config file: %v", err)
		return err
	}

	return nil
}
