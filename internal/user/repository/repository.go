package repository

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, exception.Exception)
	FindUserAndPwdByEmail(ctx context.Context, email string) (*dto.UserWithPwdDTO, exception.Exception)
}
