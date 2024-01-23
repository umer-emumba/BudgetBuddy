package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email           string    `json:"email" gorm:"unique;index"`
	Password        string    `json:"-" gorm:"size:255"`
	EmailVerifiedAt time.Time `json:"email_verified_at" gorm:"default:NULL"`
	Name            string    `json:"name" gorm:"size:255"`
	ImageUrl        string    `json:"image_url" gorm:"size:255"`
}
