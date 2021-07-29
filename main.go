package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/qbarrand/arbalete/internal/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/qbarrand/arbalete/internal/config"
)

func main() {
	cfg, err := config.ParseCommandLine(os.Args[0], os.Args[1:]...)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not parse the command line")
	}

	logger, err := getLogger(*cfg, os.Stdout)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create a logger")
	}

	a, err := app.NewApp()
	if err != nil {
		log.Fatal().Err(err).Msg("Could not create an app")
	}

	logger.
		Info().
		Str("addr", cfg.Addr).
		Str("metrics_addr", cfg.MetricsAddr).
		Msg("Starting the app")

	if err = http.ListenAndServe(cfg.Addr, a.Router()); err != nil {
		log.Fatal().Err(err).Msg("Server error")
	}
}

func getLogger(cfg config.CommandLine, w io.Writer) (zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		return zerolog.Nop(), fmt.Errorf("%q: invalid log level: %v", cfg.LogLevel, err)
	}

	if cfg.LogFormat == "json" {
		w = &zerolog.ConsoleWriter{Out: w}
	}

	logger := zerolog.New(w).Level(level).With().Timestamp().Logger()

	return logger, nil
}
