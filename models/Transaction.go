package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount            float64
	TransactionDate   time.Time
	TransactionTypeID int
	TransactionType   TransactionType
	CategoryID        int
	Category          Category
	UserID            int
	User              User
}
