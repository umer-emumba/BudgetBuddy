package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name              string `gorm:"size:255"`
	TransactionTypeID int
	TransactionType   TransactionType
}
