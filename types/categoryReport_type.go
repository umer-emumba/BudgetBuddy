package types

type CategoryReport struct {
	Category        string  `json:"category"`
	TransactionType string  `json:"transaction_type"`
	TotalAmount     float64 `json:"total_amount"`
}
