package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog"
)

func ReadEnvConfig(cfg any, opts *env.Options) error {
	return ReadEnvConfigWithLogger(zerolog.DefaultContextLogger, cfg, opts)
}

func ReadEnvConfigWithLogger(logger *zerolog.Logger, cfg any, opts *env.Options) error {
	logger.Debug().Msgf("config input: %++v", cfg)

	logger.Trace().Msgf("options: %++v", opts)

	if opts == nil {
		if err := env.Parse(cfg); err != nil {
			logger.Error().Err(err).Msg("Failed to read configuration from environment")

			return err
		}

		logger.Trace().Msgf("cfg output: %++v", cfg)

		return nil
	}

	if err := env.ParseWithOptions(cfg, *opts); err != nil {
		logger.Error().Err(err).Msg("Failed to read configuration from environment")

		return err
	}

	logger.Trace().Msgf("cfg output: %++v", cfg)

	return nil
}
