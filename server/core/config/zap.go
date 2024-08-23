package config

import "go.uber.org/zap/zapcore"

type ZapConfig struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`                      // 级别
	Prefix      string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                   // 日志前缀
	Format      string `mapstructure:"format" json:"format" yaml:"format"`                   // 输出
	EncodeLevel string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"` // 编码级
	ShowLine    bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`          // 显示行
}

// Levels returns the configured log levels.
func (cfg *ZapConfig) Levels() []zapcore.Level {
	levels := make([]zapcore.Level, 0, 1)
	level, err := zapcore.ParseLevel(cfg.Level)
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
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	if cfg.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)
}

func (cfg *ZapConfig) LevelEncoder() zapcore.LevelEncoder {
	switch {
	case cfg.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case cfg.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}
