package config

import (
	"github.com/spf13/viper"

	"github.com/sudora1n/google-auth-bot/internal/microservice-api/logger"
)

var Config MainConfig

type MainConfig struct {
	DB_Host     string `mapstructure:"DATABASE_HOST"`
	DB_Port     int    `mapstructure:"DATABASE_PORT"`
	DB_User     string `mapstructure:"DATABASE_USER"`
	DB_Password string `mapstructure:"DATABASE_PASSWORD"`
	DB_Name     string `mapstructure:"DATABASE_NAME"`
}

func InitConfig() {
	ParseFlags()
	viper.SetConfigFile(*configPath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Panicf("Fatal error config file %s", err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		logger.Logger.Panicf("Fatal error config file %s", err)
	}
}
