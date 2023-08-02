package logger

import (
	config "github.com/serhiq/effeInventory/services/inventory/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var SugaredLogger *zap.SugaredLogger

func InitLogger(cfg config.Config) (err error) {
	loggingLevel := zap.DebugLevel

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore))

	SugaredLogger = notSugaredLogger.Sugar().With(
		"service", cfg.Project.ServiceName,
	)
	return nil
}

func Sync() {
	SugaredLogger.Info("syncing zap logger")
	_ = SugaredLogger.Sync()
}
