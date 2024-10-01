package controllers

import (
	"github.com/gin-gonic/gin"
	"smolurl/database"
	"smolurl/models"
)

func UpdateAnalytics(shortUrl string) {
	var url models.URL
	if err := database.DB.Where("url = ?", shortUrl).First(&url).Error; err == nil {
		database.DB.Model(&url).UpdateColumn("clicks", url.Clicks+1)
	}
}

func GetURLStats(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	urlID := c.Param("id")

	var url models.URL

	if err := database.DB.Where("id = ? AND user_id = ?", urlID, userID).First(&url).Error; err != nil {
		c.JSON(404, gin.H{"error": "Url not found"})
		return
	}

	c.JSON(200, gin.H{
		"original_url": url.OriginalUrl,
		"short_url":    url.ShortUrl,
		"clicks":       url.ShortUrl,
	})
}
