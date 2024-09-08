package config

type ZbxConfig struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Token    string `mapstructure:"token" json:"token" yaml:"token"`
}
