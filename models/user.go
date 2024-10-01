package models

import "gorm.io/gorm"

type USER struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	Urls     []URL  `gorm:"foreignKey:UserID"`
}
