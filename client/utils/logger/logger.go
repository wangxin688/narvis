package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SetUpLogger() {
	levels := make([]zapcore.Level, 0, 1)
	level, err := zapcore.ParseLevel("info")
	if err != nil {
		level = zapcore.DebugLevel
	}
	levels = append(levels, level)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		NameKey:        "name",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(log.Writer()), levels[0])

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(Logger)
	Logger.Info("[LoggerSetup] logger set up successfully")
}
