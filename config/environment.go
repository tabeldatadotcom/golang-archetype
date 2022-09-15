package config

import (
	"github.com/spf13/viper"
	"log"
)

func SetupEnvironment() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.SetDefault("LOG_FILE_LOCATION", "./logs/application.log")
	
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
