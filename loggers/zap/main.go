package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	p1     Person
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func init() {
	p1 = Person{
		Name:  "Fernando",
		Age:   43,
		Email: "fgiorgetti@gmail.com",
	}
}

func zapConfig(encoding string) *zap.Logger {
	zapCfg := zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapCfg.Encoding = encoding
	zapCfg.Level, _ = zap.ParseAtomicLevel("info")
	logger, _ = zapCfg.Build(zap.AddCaller())
	return logger
}

func main() {
	consoleLogging()
	jsonLogging()
}

func jsonLogging() {
	logger = zapConfig("json")
	logStuff(logger)
}

func consoleLogging() {
	logger = zapConfig("console")
	logStuff(logger)
}

func logStuff(logger *zap.Logger) {
	logger.Info("info from zap")
	logger.Warn("warn from zap")
	// switch level to debug for this message to be logged
	logger.Debug("debug from zap")

	// sugared logger allows you to log your own structures (but it is not so fast)
	sl := logger.Sugar()
	sl.With("Person", p1).Info("Creating person approach 1")
	sl.Infow("Creating person approach 2", "Person", p1)

	err := fmt.Errorf("unable to write to disk")
	sl.Error(err)
}
