package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/wangxin688/narvis/intend/logger"
)

var Settings *Config

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

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Attempt to read the config file
	if err = viper.ReadInConfig(); err != nil {
		// logger.Errorf("Error reading config file, %s", zap.Error(err)
		return err
	}

	// Unmarshal the config file values into the Config struct
	if err = viper.Unmarshal(&settings); err != nil {
		return fmt.Errorf("error unmarshaling viper config into `Settings`: %w", err)
	}

	Settings = settings

	return nil
}

func SetUpLogger() {
	logConfig := logger.LogConfig{Formatter: "text"}
	logger.InitLogger(&logConfig)
}
