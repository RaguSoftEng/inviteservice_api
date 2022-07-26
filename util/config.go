package util

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Databse  string `mapstructure:"database"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Db           DbConfig `mapstructure:"db"`
	JwtSecretKey string   `mapstructure:"jwt_secretKey"`
}

func loadVariables() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

func InitVariables() {
	AppConfig = loadVariables()
}
