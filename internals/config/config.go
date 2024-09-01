package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort, ApiVersion, ApiTitle, ApiDescription string
	RedisPort, RedisHost, RedisTtl string
	Dsn, DbHost, DbName, DbUser string
	DbPort, DbPassword, JwtSecret string
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Cannot load environment file: ", err)
		return nil, err
	}

	config := &AppConfig{
		AppPort: os.Getenv("APP_PORT"),
		ApiVersion: os.Getenv("API_VERSION"),
		ApiTitle: os.Getenv("API_TITLE"),
		ApiDescription: os.Getenv("API_DESCRIPTION"),
		DbHost: os.Getenv("HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("USERNAME"),
		DbPassword: os.Getenv("PASSWORD"),
		Dsn: os.Getenv("DATABASE_PSQL_URL"),
		RedisPort: os.Getenv("REDIS_PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisTtl: os.Getenv("REDIS_TTL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}

	return config, nil
}