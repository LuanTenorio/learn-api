package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, error)
}
