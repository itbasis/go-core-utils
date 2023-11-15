package envreader

import (
	"context"

	"github.com/caarlos0/env/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"

	"github.com/juju/zaputil/zapctx"
)

func ReadEnvConfig(ctx context.Context, cfg any, opts *env.Options) error {
	sugaredLogger := zapctx.Logger(ctx).Sugar()

	if sugaredLogger.Level() == zapcore.DebugLevel {
		sugaredLogger.Debugf("config input: %++v", cfg)
		sugaredLogger.Debugf("options: %++v", opts)
	}

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
