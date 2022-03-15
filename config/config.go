package config

import "github.com/spf13/viper"

func GetConfig() error {
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}
