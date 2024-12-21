package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger(logLevel string) {
	// 设置日志级别
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	config := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	sugarLogger = logger.Sugar()
}

func Info(message string, fields ...interface{}) {
	sugarLogger.Infow(message, fields...)
}

func Warn(message string, fields ...interface{}) {
	sugarLogger.Warnw(message, fields...)
}

func Error(message string, fields ...interface{}) {
	sugarLogger.Errorw(message, fields...)
}

func Sync() {
	err := sugarLogger.Sync()
	if err != nil {
		return
	}
}
