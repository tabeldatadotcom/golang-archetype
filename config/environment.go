package config

import (
	"github.com/spf13/viper"
	"log"
)

func SetupEnvironment() {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}