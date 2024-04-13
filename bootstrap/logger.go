package bootstrap

import (
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
)

func newLogger() *zap.SugaredLogger {
	logger := rkentry.GlobalAppCtx.GetLoggerEntry("zap-logger")
	defer logger.Sync()

	return logger.Sugar()
}
