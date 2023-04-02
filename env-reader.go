package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/rs/zerolog"
)

func ReadEnvConfig(cfg any, opts ...env.Options) error {
	return ReadEnvConfigWithLogger(zerolog.DefaultContextLogger, cfg, opts...)
}

func ReadEnvConfigWithLogger(logger *zerolog.Logger, cfg any, opts ...env.Options) error {
	logger.Debug().Msgf("config input: %++v", cfg)

	if err := env.Parse(cfg, opts...); err != nil {
		logger.Error().Err(err).Msg("Failed to read configuration from environment")

		return err
	}

	logger.Trace().Msgf("cfg output: %++v", cfg)

	return nil
}
