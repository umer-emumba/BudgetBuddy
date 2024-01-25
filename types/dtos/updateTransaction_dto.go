package dtos

type UpdateTransactionDto struct {
	Amount            float64 `json:"amount" form:"amount" binding:"omitempty,numeric,min=1"`
	TransactionTypeID int     `json:"transaction_type_id" form:"transaction_type_id" binding:"omitempty"`
	CategoryID        int     `json:"category_id" form:"category_id" binding:"omitempty"`
	TransactionDate   string  `json:"transaction_date" form:"transaction_date" binding:"omitempty,datetime=2006-01-02T15:04:05"`
}
