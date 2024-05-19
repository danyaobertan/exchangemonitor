package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() *Configuration {
	viper.SetConfigName("local")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}
	Config := &Configuration{}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalf("Failed to unmarshal configuration: %s", err)
	}

	return Config
}
