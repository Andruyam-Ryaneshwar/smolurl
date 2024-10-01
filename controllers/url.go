package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"net/http"
	"smolurl/cache"
	"smolurl/database"
	"smolurl/models"
)

func ShortenURL(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		OriginalUrl string `json:"original_url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	shortCode, err := shortid.Generate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short code"})
		return
	}

	url := models.URL{
		OriginalUrl: input.OriginalUrl,
		ShortUrl:    shortCode,
		UserID:      userID,
	}

	if err := database.DB.Create(&url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save url"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "https://smolurl.com/" + shortCode})
}

func Redirect(c *gin.Context) {
	shortUrl := c.Param("url")

	cachedUrl, err := cache.GetURL(shortUrl)

	if err == nil {
		c.Redirect(http.StatusFound, cachedUrl)
		return
	}

	var url models.URL

	if err := database.DB.Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Url not found"})
		return
	}

	cache.SetURL(shortUrl, url.OriginalUrl)
}
