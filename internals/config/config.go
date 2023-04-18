package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MySQLHost           string `mapstructure:"MYSQL_HOST"`
	MySQLPort           string `mapstructure:"MYSQL_PORT"`
	MySQLDatabase       string `mapstructure:"MYSQL_DATABASE"`
	MySQLUser           string `mapstructure:"MYSQL_USER"`
	MySQLPassword       string `mapstructure:"MYSQL_PASSWORD"`
	RedisHost           string `mapstructure:"REDIS_HOST"`
	RedisPort           string `mapstructure:"REDIS_PORT"`
	SessionMaxAgeInDays int    `mapstructure:"SESSION_MAX_AGE_IN_DAYS"`
	SessionSecret       string `mapstructure:"SESSION_SECRET"`
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return
}
