package core

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
	"github.com/wangxin688/narvis/server/core/config"
)

var Settings *config.Settings

var ProjectPath string

var Environment config.Env

// setup viper config from config.yaml and make Settings as a global variable
func SetUpConfig() {
	var settings config.Settings

	_, currentPath, _, _ := runtime.Caller(0)

	projectPath := filepath.Dir(filepath.Dir(currentPath))

	ProjectPath = projectPath

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(projectPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %s \n", err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Settings = &settings

	Environment = settings.Env
}
