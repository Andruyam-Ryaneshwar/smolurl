package database

import (
	"log"
	"smolurl/config"
	"smolurl/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(config.AppConfig.DB_URL), &gorm.Config{})

	if err != nil {
		log.Println("Error while creating DB instance")
	}

	err = db.AutoMigrate(&models.USER{}, &models.URL{})

	if err != nil {
		log.Println("Failed to migrate database")
	}

	DB = db
}
