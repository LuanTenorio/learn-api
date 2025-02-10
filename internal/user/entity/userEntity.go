package entity

import "github.com/LuanTenorio/learn-api/internal/user/dto"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserByCreateDto(u *dto.CreateUserDTO, id int) *User {
	return &User{
		Id:    id,
		Name:  u.Name,
		Email: u.Email,
	}
}
