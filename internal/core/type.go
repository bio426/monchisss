package core

type Controller struct{}

type Service struct{}

type PaginatedResponse[T any] struct {
	Total int32 `json:"total"`
	From  int32 `json:"from"`
	To    int32 `json:"to"`
	Rows  []T   `json:"rows"`
}
