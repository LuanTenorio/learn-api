package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *dto.CreateUserDTO) (int, error)
	FindUserAndPwdByEmail(ctx context.Context, email string) (*dto.UserWithPwdDTO, error)
}
