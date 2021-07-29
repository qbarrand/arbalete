package config_test

import (
	"testing"

	"github.com/qbarrand/arbalete/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCommandLine(t *testing.T) {
	const (
		addr        = ":1234"
		logFormat   = "some-format"
		logLevel    = "debug"
		metricsAddr = ":5678"
	)

	args := []string{
		"--addr", addr,
		"--log-format", logFormat,
		"--log-level", logLevel,
		"--metrics-addr", metricsAddr,
	}

	expected := &config.CommandLine{
		Addr:        addr,
		LogFormat:   logFormat,
		LogLevel:    logLevel,
		MetricsAddr: metricsAddr,
	}

	cfg, err := config.ParseCommandLine("", args...)

	require.NoError(t, err)
	assert.Equal(t, expected, cfg)
}
