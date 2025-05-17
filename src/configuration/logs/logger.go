package logs

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log         *zap.Logger
	OUTPUT_LOGS = "OUTPUT_LOGS"
	LEVEL_LOGS  = "LEVEL_LOGS"
)

func init() {
	logConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(GetLevelLogs()),
		Development: true,
		Encoding:    "json",
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
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{GetOutputLogs()},
	}

	log, _ = logConfig.Build()
}

func GetOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(OUTPUT_LOGS)))

	if output == "" {
		return "stdout"
	} else {
		return output
	}
}

func GetLevelLogs() zapcore.Level {
	level := strings.ToLower(strings.TrimSpace(os.Getenv(LEVEL_LOGS)))

	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}
