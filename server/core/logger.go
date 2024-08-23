package core

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SetUpLogger() {
	level := Settings.Zap.Levels()
	encoder := Settings.Zap.Encoder()

	core := zapcore.NewCore(encoder, zapcore.AddSync(log.Writer()), level[0])

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(Logger)
}
