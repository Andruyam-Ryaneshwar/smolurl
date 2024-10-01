package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SERVER_PORT string
	DB_URL      string
	JWT_SECRET  string
	RedisAddr   string
	RedisPwd    string
}

var AppConfig *Config

func InitConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Println("Error while initalizing config")
	}

	AppConfig = &Config{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		DB_URL:      os.Getenv("DB_URL"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		RedisAddr:   os.Getenv("RedisAddr"),
		RedisPwd:    os.Getenv("RedisPwd"),
	}
}
