package configs

import (
	"log"

	"github.com/spf13/viper"
)

type ENVIRONMENTS string;

const (
	DEVELOPMENT ENVIRONMENTS = "development"
	PRODUCTION ENVIRONMENTS = "production"
) 
type Configs struct {
	PORT string 
	ENVIRONMENT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	DB_HOST string
	DB_PORT string
	DB_URL string
	REDIS_URL string
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

	var configs Configs;

	viper.Unmarshal(&configs)

	return configs
}