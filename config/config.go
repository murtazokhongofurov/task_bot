package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)


type Config struct {
	BotToken string 
	SqliteUrl string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	
	conf.AutomaticEnv()

	cfg := Config{
		BotToken: conf.GetString("BOT_TOKEN"),
		SqliteUrl: conf.GetString("SQLITE_URL"),
	}
	return cfg
}