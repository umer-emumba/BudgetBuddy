package types

type Pagination[Model any] struct {
	Count int     `json:"count"`
	Rows  []Model `json:"rows"`
}
