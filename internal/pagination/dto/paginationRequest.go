package dto

type PaginationRequestDTO struct {
	Page  int `query:"page" validate:"number"`
	Limit int `query:"limit" validate:"number,max=50"`
}
