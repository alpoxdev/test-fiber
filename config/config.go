package config

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type Config struct {
	GO_ENV string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_MAX_CONNECTION int
}

var AppConfig Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	log.Info(os.Getenv("DB_MAX_CONNECTION"))

	AppConfig = Config{
		GO_ENV:            os.Getenv("GO_ENV"),
		DB_HOST:           os.Getenv("DB_HOST"),
		DB_USER:           os.Getenv("DB_USER"),
		DB_PASSWORD:       os.Getenv("DB_PASSWORD"),
		DB_NAME:           os.Getenv("DB_NAME"),
		DB_PORT:           os.Getenv("DB_PORT"),
		DB_MAX_CONNECTION: func() int {
			if val, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION")); err == nil {
				return val
			}
			return 1
		}(),
	}

	log.Info("Config loaded successfully")
}
