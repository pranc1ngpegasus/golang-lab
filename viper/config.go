package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var _ Configuration = (*config)(nil)

type Configuration interface {
	Debug() *Debug
	Server() *Server
}

type (
	Config struct {
		Debug  Debug
		Server Server
	}

	Debug struct {
		Enabled      bool
		LogFrequency time.Duration
	}

	Server struct {
		Port string
	}
)

type config struct {
	debug  *Debug
	server *Server
}

func NewConfig() (Configuration, error) {
	viper.SetConfigFile("sample.yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &config{
		debug:  &cfg.Debug,
		server: &cfg.Server,
	}, nil
}

func (c *config) Debug() *Debug {
	return c.debug
}

func (c *config) Server() *Server {
	return c.server
}
