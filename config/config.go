package config

import (
	"github.com/spf13/viper"
)

// Read config in /config dir
func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
