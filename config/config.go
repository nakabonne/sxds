package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config is config for sxds
	Config struct {
		Env     string `envconfig:"ENV" default:"develop"`
		AdsMode bool   `envconfig:"ADS_MODE" default:"false"`

		Xds
		Cacher
	}

	// Xds is config for xDS server
	Xds struct {
		Port int `envconfig:"XDS_PORT" default:"8081"`
	}

	// Cacher is config for cacher server
	Cacher struct {
		Port int `envconfig:"CACHER_PORT" default:"8082"`
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
