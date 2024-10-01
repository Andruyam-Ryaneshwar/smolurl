package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	OriginalUrl string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
	Clicks      uint   `json:"clicks"`
	UserID      uint   `json:"user_id"`
}
