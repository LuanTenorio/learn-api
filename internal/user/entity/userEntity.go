package entity

import (
	"github.com/LuanTenorio/learn-api/internal/database/sqlc"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
)

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

func M2E(u *sqlc.CreateUserRow) *User {
	return &User{
		Id:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time.String(),
	}
}
