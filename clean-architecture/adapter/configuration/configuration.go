package configuration

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Server Server
	}

	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"3000"`
	}
)

var globalConfig Config

func Get() (Config, error) {
	var emptyConfig Config
	if globalConfig == emptyConfig {
		if err := envconfig.Process("", &globalConfig); err != nil {
			return Config{}, fmt.Errorf("failed to load environment variables: %w", err)
		}
	}

	return globalConfig, nil
}