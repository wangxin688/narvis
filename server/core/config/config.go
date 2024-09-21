package config

import "fmt"

type Env string

const (
	Dev   Env = "dev"
	Stage Env = "stage"
	Prod  Env = "prod"
)

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
	RabbitMQ RabbitMQConfig `mapstructure:"rmq" json:"rmq" yaml:"rmq"`
}

// alertManager config
type AtmConfig struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type CORS struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers" json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
}

type JwtConfig struct {
	SecretKey                 string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Algorithm                 string `mapstructure:"algorithm" json:"algorithm" yaml:"algorithm"`
	AccessTokenExpiredMinute  int    `mapstructure:"access_token_expired_minutes" json:"access_token_expired_minutes" yaml:"access_token_expired_minutes"`
	RefreshTokenExpiredMinute int    `mapstructure:"refresh_token_expired_minutes" json:"refresh_token_expired_minutes" yaml:"refresh_token_expired_minutes"`
}

type PostgresConfig struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DBName       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode      string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
}

func (cfg PostgresConfig) BuildPgDsn() string {
	return "host=" + cfg.Host + " user=" + cfg.Username + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}

type RedisConfig struct {
	Host        string   `mapstructure:"host" json:"host" yaml:"host"`
	Port        int      `mapstructure:"port" json:"port" yaml:"port"`
	Password    string   `mapstructure:"password" json:"password" yaml:"password"`
	Db          int      `mapstructure:"db" json:"db" yaml:"db"`
	UseCluster  bool     `mapstructure:"use_cluster" json:"use_cluster" yaml:"use_cluster"`
	ClusterAddr []string `mapstructure:"cluster_addr" json:"cluster_addr" yaml:"cluster_addr"`
}

func (rd RedisConfig) BuildRedisDsn() string {
	return fmt.Sprintf("redis://:%s:@%s:%d/%d", rd.Password, rd.Host, rd.Port, rd.Db)
}

type SentryConfig struct {
	Dsn             string  `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	EnableTracing   bool    `mapstructure:"enable_tracing" json:"enable_tracing" yaml:"enable_tracing"`
	TraceSampleRate float64 `mapstructure:"trace_sample_rate" json:"trace_sample_rate" yaml:"trace_sample_rate"`
	Release         string  `mapstructure:"release" json:"release" yaml:"release"`
}

type SystemConfig struct {
	RouterPrefix string `mapstructure:"router_prefix" json:"router_prefix" yaml:"router_prefix"`
	BaseUrl      string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
}

type VtmConfig struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type ZbxConfig struct {
	Url   string `mapstructure:"url" json:"url" yaml:"url"`
	Token string `mapstructure:"token" json:"token" yaml:"token"`
}

type RabbitMQConfig struct {
	ServerUrl string `mapstructure:"server_url" json:"server_url" yaml:"server_url"`
	ProxyUrl  string `mapstructure:"proxy_url" json:"proxy_url" yaml:"proxy_url"`
}
