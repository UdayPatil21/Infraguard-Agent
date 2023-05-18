package configHelper

import (
	"infraguard-agent/helpers/logger"

	"github.com/spf13/viper"
)

const (
	ConfigDirPath  = "../config"
	ConfigFilePath = "../config/config.json"
)

var cnf *viper.Viper

// Load config file from given path
func InitConfig() {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigFile(ConfigFilePath)
	v.AddConfigPath(ConfigDirPath)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error("config file not found", err)
		}
		logger.Error("config file not found", err)
		logger.Info("config file not found", err)
	}

	cnf = v
}

func GetString(key string) string {
	return cnf.GetString(key)
}
