package config

type JwtConfig struct {
	SecretKey                 string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Algorithm                 string `mapstructure:"algorithm" json:"algorithm" yaml:"algorithm"`
	AccessTokenExpiredMinute  int    `mapstructure:"access_token_expired_minute" json:"access_token_expired_minute" yaml:"access_token_expired_minute"`
	RefreshTokenExpiredMinute int    `mapstructure:"refresh_token_expired_minute" json:"refresh_token_expired_minute" yaml:"refresh_token_expired_minute"`
}
