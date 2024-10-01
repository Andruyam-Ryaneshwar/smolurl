package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smolurl/utils"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		userID, err := utils.ParseToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		c.Set("userID", userID)
		c.Next()
	}
}
