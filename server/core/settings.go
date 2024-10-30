package core

import (
	"log"

	"github.com/spf13/viper"
	"github.com/wangxin688/narvis/server/core/config"
)

var Settings *config.Settings

// setup viper config from config.yaml and make Settings as a global variable
func SetUpConfig() {
	var settings config.Settings

	// _, currentPath, _, _ := runtime.Caller(0)

	// projectPath := filepath.Dir(filepath.Dir(currentPath))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s \n", err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Settings = &settings
}
