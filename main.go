package main

import (
	"log"
	// "smolurl/cache"
	"smolurl/config"
	"smolurl/database"
	"smolurl/routes"
)

func main() {
	config.InitConfig()

	database.ConnectDB()

	// cache.InitRedis()

	r := routes.SetupRouter()

	err := r.Run(":" + config.AppConfig.SERVER_PORT)

	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
