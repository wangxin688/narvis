package config

type VtmConfig struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
