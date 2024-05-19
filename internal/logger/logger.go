package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

func InitLogger(debug bool) *zap.Logger {
	encoderConfig := zap.NewProductionConfig()
	if debug {
		encoderConfig.Level.SetLevel(zap.DebugLevel)
	}

	encoderConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := encoderConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("Error initializing logger: %s", err))
	}

	_ = logger.Sync()

	zap.ReplaceGlobals(logger)

	return logger
}
