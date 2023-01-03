package config

import "github.com/Ayush-Walia/Fampay-Youtube/utils"

type AppConfig struct {
	ServerPort string
	DBUser     string
	DBPass     string
	DBAddr     string
	DBName     string
}

func LoadConfig() *AppConfig {
	appConfig := new(AppConfig)
	appConfig.ServerPort = utils.GetEnv("ServerPort", "8080")

	return appConfig
}
