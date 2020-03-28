package config

import (
	"os"
)

// Config contains the environment specific configuration values needed by the
// application.
type Config struct {
	Environment string
	Port        int
}

const environmentEnv = "ENVIRONMENT"

// New returns an instance of Config based on the "ENVIRONMENT" environment
// variable.
func New() Config {
	cfg := Config{
		Port: 3000,
	}

	switch os.Getenv(environmentEnv) {
	case "development", "":
		loadDevelopmentConfig(&cfg)
	case "test":
		loadTestConfig(&cfg)
	}

	return cfg
}
