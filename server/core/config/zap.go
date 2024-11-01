package config

import "go.uber.org/zap/zapcore"

type ZapConfig struct { // 级别
	Format string `mapstructure:"format" json:"format" yaml:"format"` // 输出
	// EncodeLevel string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"` // 编码级
}

// Levels returns the configured log levels.
func (cfg *ZapConfig) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 1)
	level, err := zapcore.ParseLevel("info")
	if err != nil {
		level = zapcore.DebugLevel
	}
	levels = append(levels, level)
	return levels
}

func (cfg *ZapConfig) Encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
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
	if cfg.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

