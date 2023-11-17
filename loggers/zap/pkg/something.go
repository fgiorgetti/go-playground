package pkg

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.Level, _ = zap.ParseAtomicLevel("info")
	logger, _ = zapCfg.Build()
}

func Something() {
	logger.Info("Something info")
	logger.Debug("Something debug")
}

func SomethingElse() {
	logger.Info("SomethingElse info")
	logger.Debug("SomethingElse debug")
}
