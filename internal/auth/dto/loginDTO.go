package dto

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email,max=60"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}
