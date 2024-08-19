package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Api_Version, Api_Title, Api_Description string
	Redis_Port, Redis_Host, Redis_Ttl string
	Db_Host, Db_Name, Db_User string
	Db_Port, Db_Password string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Cannot load environment file: ", err)
		return nil, err
	}

	config := &Config{
		Api_Version: os.Getenv("API_VERSION"),
		Api_Title: os.Getenv("API_TITLE"),
		Api_Description: os.Getenv("API_DESCRIPTION"),
		Db_Host: os.Getenv("HOST"),
		Db_Name: os.Getenv("DB_NAME"),
		Db_Port: os.Getenv("DB_PORT"),
		Db_User: os.Getenv("USERNAME"),
		Db_Password: os.Getenv("PASSWORD"),
	}

	return config, nil
}