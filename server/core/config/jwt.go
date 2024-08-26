package config

type JwtConfig struct {
	SecretKey                 string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Algorithm                 string `mapstructure:"algorithm" json:"algorithm" yaml:"algorithm"`
	AccessTokenExpiredMinute  int    `mapstructure:"access_token_expired_minutes" json:"access_token_expired_minutes" yaml:"access_token_expired_minutes"`
	RefreshTokenExpiredMinute int    `mapstructure:"refresh_token_expired_minutes" json:"refresh_token_expired_minutes" yaml:"refresh_token_expired_minutes"`
}
