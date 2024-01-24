package dtos

type CreateTransactionDto struct {
	Amount            float64 `json:"amount" form:"amount" binding:"required,min=1"`
	TransactionTypeID int     `json:"transaction_type_id" form:"transaction_type_id" binding:"required"`
	CategoryID        int     `json:"category_id" form:"category_id" binding:"required"`
	TransactionDate   string  `json:"transaction_date" form:"transaction_date" binding:"required,datetime=2006-01-02T15:04:05"`
}
