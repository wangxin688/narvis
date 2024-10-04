package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var Settings *Config

var ProjectPath string

type Environment string

const (
	Dev   Environment = "dev"
	Stage Environment = "stage"
	Prod  Environment = "prod"
)

type Config struct {
	ORGANIZATION_ID string `mapstructure:"ORGANIZATION_ID"`
	PROXY_ID        string `mapstructure:"PROXY_ID"`
	SECRET_KEY      string `mapstructure:"SECRET_KEY"`
	SERVER_URL      string `mapstructure:"SERVER_URL"`
	AMQP_URL        string `mapstructure:"AMQP_URL"`
}

func (c *Config) WebSocketUrl() string {
	if strings.Contains(c.SERVER_URL, "https") {
		return strings.Replace(c.SERVER_URL, "https", "wss", 1)
	} else if strings.Contains(c.SERVER_URL, "http") {
		return strings.Replace(c.SERVER_URL, "http", "ws", 1)
	}
	return c.SERVER_URL
}

// SetupConfig sets up the Config struct from the .env file in the root of the project.
// If the .env file is not found, it will set defaults for the fields.
// The .env file must be in KEY=VALUE format. The fields of the Config struct must
// match the keys in the .env file.
//
// It returns an error if there is an issue reading the file or unmarshaling the
// config into the Config struct.
func SetupConfig() (err error) {

	var settings = Settings

	_, currentPath, _, _ := runtime.Caller(0)

	projectPath := filepath.Dir(filepath.Dir(currentPath))

	// Set defaults for the config file values
	viper.SetDefault("AccessTokenExpireMinutes", time.Duration(60*24)*time.Minute)
	viper.SetDefault("RefreshTokenExpireMinutes", time.Duration(60*24*8)*time.Minute)
	viper.SetDefault("ServerTimeZone", "Asia/Shanghai")
	viper.SetDefault("DBDebugMode", false)
	viper.SetDefault("RedisPort", 6379)
	viper.SetDefault("RedisDB", 0)

	// Set the config file name and type
	viper.SetConfigFile(projectPath + "/.env")
	viper.SetConfigType("env")

	// Attempt to read the config file
	if err = viper.ReadInConfig(); err != nil {
		// logger.Errorf("Error reading config file, %s", err)
		return err
	}

	// Unmarshal the config file values into the Config struct
	if err = viper.Unmarshal(&settings); err != nil {
		return fmt.Errorf("error unmarshaling viper config into `Settings`: %w", err)
	}

	// Update the ProjectPath var
	ProjectPath = projectPath

	Settings = settings

	return nil
}
