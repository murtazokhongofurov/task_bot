package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)


type Config struct {
	BotToken string 
	SqliteUrl string
}

func Load() Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	c := Config{}

	c.BotToken = cast.ToString(getOrReturnDefault("BOT_TOKEN", "bot_token"))
	c.SqliteUrl = cast.ToString(getOrReturnDefault("SQLITE_URL", "sqlite_url"))
	return c
}


func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}