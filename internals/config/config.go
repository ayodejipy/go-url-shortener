package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ApiVersion, ApiTitle, ApiDescription string
	RedisPort, RedisHost, RedisTtl string
	DbHost, DbName, DbUser string
	DbPort, DbPassword string
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Cannot load environment file: ", err)
		return nil, err
	}

	config := &AppConfig{
		ApiVersion: os.Getenv("API_VERSION"),
		ApiTitle: os.Getenv("API_TITLE"),
		ApiDescription: os.Getenv("API_DESCRIPTION"),
		DbHost: os.Getenv("HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("USERNAME"),
		DbPassword: os.Getenv("PASSWORD"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisTtl: os.Getenv("REDIS_TTL"),
	}

	return config, nil
}