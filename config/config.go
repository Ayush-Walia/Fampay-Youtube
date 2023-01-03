package config

import (
	"os"

	"github.com/Ayush-Walia/Fampay-Youtube/utils"
	"github.com/gookit/slog"
)

type AppConfig struct {
	ServerPort   string
	DBUser       string
	DBPass       string
	DBAddr       string
	DBName       string
	APIKeys      string
	YoutubeQuery string
	YoutubeCron  string
}

func LoadConfig() *AppConfig {
	appConfig := new(AppConfig)
	appConfig.ServerPort = utils.GetEnv("ServerPort", "8080")
	appConfig.DBUser = utils.GetEnv("DBUser", "root")
	appConfig.DBPass = utils.GetEnv("DBPass", "mysql123")
	appConfig.DBAddr = utils.GetEnv("DBAddr", "localhost:3306")
	appConfig.DBName = utils.GetEnv("DBName", "fampay_youtube")
	appConfig.YoutubeQuery = utils.GetEnv("YoutubeQuery", "football")
	appConfig.YoutubeCron = utils.GetEnv("YoutubeCron", "@every 10m")
	appConfig.APIKeys = os.Getenv("APIKeys")

	if len(appConfig.APIKeys) == 0 {
		slog.Fatal("APIKeys missing, please pass APIKeys env variable")
	}

	return appConfig
}
