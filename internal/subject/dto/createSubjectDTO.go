package dto

type CreateSubjectDTO struct {
	Name string `json:"name" validate:"required,min=3,max=150"`
}
