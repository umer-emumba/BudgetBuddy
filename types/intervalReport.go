package types

type IntervalReport struct {
	Interval        string  `json:"interval"`
	TransactionType string  `json:"transaction_type"`
	TotalAmount     float64 `json:"total_amount"`
}
