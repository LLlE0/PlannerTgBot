package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Error initializing config file!")
	}
}

// handle config file
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
