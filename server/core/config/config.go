package config

type Settings struct {
	Zap      ZapConfig      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Postgres PostgresConfig `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	Jwt      JwtConfig      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Cors     CORS           `mapstructure:"cors" json:"cors" yaml:"cors"`
	System   SystemConfig   `mapstructure:"sys" json:"sys" yaml:"sys"`
	Sentry   SentryConfig   `mapstructure:"sentry" json:"sentry" yaml:"sentry"`
	Env      Env            `mapstructure:"env" json:"env" yaml:"env"`
	Zbx      ZbxConfig      `mapstructure:"zbx" json:"zbx" yaml:"zbx"`
	Vtm      VtmConfig      `mapstructure:"vtm" json:"vtm" yaml:"vtm"`
	Atm      AtmConfig      `mapstructure:"atm" json:"atm" yaml:"atm"`
}
