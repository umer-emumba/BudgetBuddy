package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `gorm:"unique;index"`
	Password        string    `gorm:"size:255"`
	EmailVerifiedAt time.Time `gorm:"default:NULL"`
	Name            string    `gorm:"size:255"`
	ImageUrl        string    `gorm:"size:255"`
}
