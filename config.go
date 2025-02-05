package gotel

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Config holds the configuration for the telemetry.
type Config struct {
	ServiceName    string `env:"SERVICE_NAME"      envDefault:"gotel"`
	ServiceVersion string `env:"SERVICE_VERSION"   envDefault:"0.0.1"`
	Enabled        bool   `env:"TELEMETRY_ENABLED" envDefault:"true"`
}

// NewConfigFromEnv creates a new telemetry config from the environment.
func NewConfigFromEnv() (Config, error) {
	telem := Config{}
	if err := env.Parse(&telem); err != nil {
		return Config{}, fmt.Errorf("failed to parse telemetry config: %w", err)
	}

	return telem, nil
}
