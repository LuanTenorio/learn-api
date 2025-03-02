package dto

type PaginationResponseDTO struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Pages      int         `json:"pages"`
	TotalItems int         `json:"total_items"`
	Data       interface{} `json:"data"`
}
