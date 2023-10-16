package config

import (
	"context"

	"github.com/pkg/errors"

	"github.com/caarlos0/env/v9"
	"github.com/juju/zaputil/zapctx"
)

func ReadEnvConfig(ctx context.Context, cfg any, opts *env.Options) error {
	sugaredLogger := zapctx.Logger(ctx).Sugar()
	sugaredLogger.Debugf("config input: %++v", cfg)

	sugaredLogger.Debugf("options: %++v", opts)

	if opts == nil {
		if err := env.Parse(cfg); err != nil {
			e := errors.Wrap(err, "failed to read configuration from environment")
			sugaredLogger.Error(e)

			return e
		}

		sugaredLogger.Debugf("cfg output: %++v", cfg)

		return nil
	}

	if err := env.ParseWithOptions(cfg, *opts); err != nil {
		e := errors.Wrap(err, "failed to read configuration from environment")
		sugaredLogger.Error(e)

		return e
	}

	sugaredLogger.Debugf("cfg output: %++v", cfg)

	return nil
}
