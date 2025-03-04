package dto

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email,max=60"`
	Password string `json:"password" validate:"required,min=8,max=70"`
}
