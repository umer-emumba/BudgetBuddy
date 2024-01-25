package types

type MonthlyReport struct {
	TransactionType string  `json:"transaction_type"`
	TotalAmount     float64 `json:"total_amount"`
}
