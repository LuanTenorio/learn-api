package dto

type CreateSubjectDTO struct {
	Name   string `json:"name" validate:"required,min=3,max=150"`
	UserId int    `json:"-" db:"user_id"`
}
