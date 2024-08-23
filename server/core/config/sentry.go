package config

type SentryConfig struct {
	Dsn             string  `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	EnableTracing   bool    `mapstructure:"enable_tracing" json:"enable_tracing" yaml:"enable_tracing"`
	TraceSampleRate float64 `mapstructure:"trace_sample_rate" json:"trace_sample_rate" yaml:"trace_sample_rate"`
	Release         string  `mapstructure:"release" json:"release" yaml:"release"`
}
