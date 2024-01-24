package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name              string           `gorm:"size:255" json:"name"`
	TransactionTypeID int              `json:"transaction_type_id"`
	TransactionType   *TransactionType `json:"transaction_type,omitempty"`
}
