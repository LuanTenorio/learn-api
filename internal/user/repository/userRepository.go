package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (int, error)
}
