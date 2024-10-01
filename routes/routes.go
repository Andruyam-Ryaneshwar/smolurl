package routes

import (
	"github.com/gin-gonic/gin"
	"smolurl/controllers"
	"smolurl/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/:url", controllers.Redirect)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/shorten", controllers.ShortenURL)
		authorized.GET("/url/:id/stats", controllers.GetURLStats)
		// TODO: to be implemented
		// authorized.GET("/user/urls", controllers.GetUserURLs)
		// authorized.PUT("/url/:id", controllers.EditUrl)
		// authorized.DELETE("/url/:id", controllers.DeleteURL)
	}

	return router
}
