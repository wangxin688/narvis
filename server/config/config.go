package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
	"github.com/wangxin688/narvis/intend/logger"
)

type Env string

const (
	Dev    Env = "dev"
	Prod   Env = "prod"
	OnPrem Env = "on_prem"
)

type Config struct {
	Logging         logger.LogConfig `mapstructure:"logging" json:"logging" yaml:"logging"`
	Postgres        PostgresConfig   `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	Jwt             JwtConfig        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis           RedisConfig      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Cors            CORS             `mapstructure:"cors" json:"cors" yaml:"cors"`
	System          SystemConfig     `mapstructure:"sys" json:"sys" yaml:"sys"`
	Sentry          SentryConfig     `mapstructure:"sentry" json:"sentry" yaml:"sentry"`
	Env             Env              `mapstructure:"env" json:"env" yaml:"env"`
	Zbx             ZbxConfig        `mapstructure:"zbx" json:"zbx" yaml:"zbx"`
	Vtm             VtmConfig        `mapstructure:"vtm" json:"vtm" yaml:"vtm"`
	Atm             AtmConfig        `mapstructure:"atm" json:"atm" yaml:"atm"`
	RabbitMQ        RabbitMQConfig   `mapstructure:"rmq" json:"rmq" yaml:"rmq"`
	BootstrapConfig BootstrapConfig  `mapstructure:"bootstrap" json:"bootstrap" yaml:"bootstrap"`
	ClickHouse      ClickHouseConfig `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
}

// alertManager config
type AtmConfig struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type CORS struct {
	AllowedOrigins []string `mapstructure:"allowed_origins" json:"allowed_origins" yaml:"allowed_origins"`
}

func (c *CORS) validate() {
	if len(c.AllowedOrigins) == 0 {
		logger.Logger.Fatal("cors.allowed_origins is required")
	}
}

type JwtConfig struct {
	SecretKey                 string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	PublicAuthKey             string `mapstructure:"public_auth_key" json:"public_auth_key" yaml:"public_auth_key"`
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
	ServerPort int    `mapstructure:"server_port" json:"server_port" yaml:"server_port"`
	BaseUrl    string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
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
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Port      int    `mapstructure:"port" json:"port" yaml:"port"`
	Username  string `mapstructure:"username" json:"username" yaml:"username"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	ServerUrl string `mapstructure:"server_url" json:"server_url" yaml:"server_url"`
	ProxyUrl  string `mapstructure:"proxy_url" json:"proxy_url" yaml:"proxy_url"`
}

type BootstrapConfig struct {
	Organization   string `mapstructure:"organization" json:"organization" yaml:"organization"`
	EnterpriseCode string `mapstructure:"enterprise_code" json:"enterprise_code" yaml:"enterprise_code"`
	DomainName     string `mapstructure:"domain_name" json:"domain_name" yaml:"domain_name"`
	AdminPassword  string `mapstructure:"admin_password" json:"admin_password" yaml:"admin_password"`
	SnmpCommunity  string `mapstructure:"snmp_community" json:"snmp_community" yaml:"snmp_community"`
	SnmpPort       uint16 `mapstructure:"snmp_port" json:"snmp_port" yaml:"snmp_port"`
	SnmpTimeout    uint8  `mapstructure:"snmp_timeout" json:"snmp_timeout" yaml:"snmp_timeout"`
	CliUser        string `mapstructure:"cli_user" json:"cli_user" yaml:"cli_user"`
	CliPassword    string `mapstructure:"cli_password" json:"cli_password" yaml:"cli_password"`
}

type ClickHouseConfig struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	DataBase    string `mapstructure:"database" json:"database" yaml:"database"`
	ReadTimeout uint8  `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	DialTime    uint8  `mapstructure:"dial_timeout" json:"dial_timeout" yaml:"dial_timeout"`
}

func (cfg ClickHouseConfig) BuildClickHouseDsn() string {
	return fmt.Sprintf(
		"clickhouse://%s:%s@%s:%s/%s?dial_timeout=%ds&read_timeout=%ds", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DataBase, cfg.DialTime, cfg.ReadTimeout)
}

var Settings *Config

func InitConfig() {
	var settings Config
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	viper.SetConfigFile(filepath.Join(path, "config.yaml"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s \n", err)
		panic(err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	settings.Cors.validate()
	Settings = &settings
}

func InitTestFixtureConfig() {
	var settings Config
	_, filepath, _, _ := runtime.Caller(0)
	configPath := strings.Split(filepath, "config/config.go")[0]
	viper.SetConfigFile(configPath + "config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s \n", err)
		panic(err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	settings.Cors.validate()
	Settings = &settings

}

func InitLogger() {
	if Settings == nil {
		InitConfig()
	}
	logger.InitLogger(&Settings.Logging)
}
