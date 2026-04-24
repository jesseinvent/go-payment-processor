package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	PORT string
	DATABASE_URL string
}

// Load Env variables and return a Configs struct

func LoadConfigs() Configs {

	log.Println("Loading config...")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	log.Println("Config loaded successfully")

	return Configs{
		PORT:         viper.GetString("PORT"),
		DATABASE_URL: viper.GetString("DATABASE_URL"),
	}
}