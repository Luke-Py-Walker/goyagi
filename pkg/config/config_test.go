package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	cfg := New()
	assert.NotNil(t, cfg, "returned config shouldn't be nil")
}

func TestEnvironments(t *testing.T) {
	originalEnv := os.Getenv(environmentEnv)
	defer func() {
		err := os.Setenv(environmentEnv, originalEnv)
		require.NoError(t, err)
	}()

	envs := []string{"development", "test"}

	for _, env := range envs {
		err := os.Setenv(environmentEnv, env)
		require.NoError(t, err)

		cfg := New()
		assert.Equal(t, cfg.Environment, env, "incorrect environment")
	}
}
