package config

import "fmt"

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
