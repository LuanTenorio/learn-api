package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/auth/dto"
	"github.com/LuanTenorio/learn-api/internal/exception"
)

type AuthUseCase interface {
	Login(ctx context.Context, loginDto *dto.LoginDTO) (string, exception.Exception)
}
