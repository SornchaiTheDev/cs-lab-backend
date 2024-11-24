package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func NewConfig() *Config {
	var config Config

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file")
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Unable to decode into struct")
	}

	return &config
}
