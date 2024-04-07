package main

import (
	"github.com/LLlE0/PlannerTgBot/types"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("Error initializing config file!")
	}

	var Handler types.Handler = *types.NewHandler()
	if err := Handler.Database.Instance.Ping(); err != nil {
		log.Print("Error establishing connection to the database!")
	} else {
		log.Print("Successfully created a handler")
	}
}

// handle config file
func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
