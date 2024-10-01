package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"smolurl/config"
)

var DB *gorm.DB

func ConnectDB() {
	DB, err := gorm.Open(postgres.Open(config.AppConfig.DB_URL), &gorm.Config{})

	if err != nil {
		log.Println("Error while creating DB instance")
	}
}
