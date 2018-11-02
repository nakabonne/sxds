package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config is config for sxds
	Config struct {
		Env string `envconfig:"ENV" default:"develop"`

		Xds
	}

	// Xds is config for xDS server
	Xds struct {
		Port int `envconfig:"XDS_PORT" default:"8081"`
	}
)

// New generates a Config
func New() (*Config, error) {
	config := Config{}
	if err := envconfig.Process("SXDS", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) IsProduction() bool {
	return c.Env == "production"
}
