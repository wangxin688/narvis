package config

type SystemConfig struct {
	RouterPrefix string `mapstructure:"router_prefix" json:"router_prefix" yaml:"router_prefix"`
	BaseUrl      string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
}
