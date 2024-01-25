package models

import (
	"time"

	"gorm.io/gorm"
)

// make relation as pointer so that we can omit them when json parsing if they are null
type Transaction struct {
	gorm.Model
	Amount            float64          `json:"amount"`
	TransactionDate   time.Time        `json:"transaction_date"`
	TransactionTypeID int              `json:"transaction_type_id"`
	TransactionType   *TransactionType `json:"transaction_type,omitempty"`
	CategoryID        int              `json:"category_id"`
	Category          *Category        `json:"category,omitempty"`
	UserID            int              `json:"user_id"`
	User              *User            `json:"user,omitempty"`
}
