package dtos

type PaginationDto struct {
	Limit             int    `json:"limit" form:"limit" binding:"required,numeric,min=1,max=50"`
	Offset            int    `json:"offset" form:"offset" binding:"omitempty,numeric,min=0"`
	Order             string `json:"order" form:"order" binding:"required,oneof=asc desc"`
	Text              string `json:"text" form:"text" binding:"omitempty"`
	TransactionTypeId int    `json:"transaction_type_id" form:"transaction_type_id" binding:"omitempty"`
}
