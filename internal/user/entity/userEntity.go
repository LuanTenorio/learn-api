package entity

import "github.com/LuanTenorio/learn-api/internal/user/dto"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func NewUserByCreateDto(u *dto.CreateUserDTO, createdAt string, id int) *User {
	return &User{
		Id:        id,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: createdAt,
	}
}
