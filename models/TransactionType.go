package models

import (
	"gorm.io/gorm"
)

type TransactionType struct {
	gorm.Model
	Name string `gorm:"size:255"`
}
