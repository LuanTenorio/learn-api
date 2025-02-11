package usecase

import (
	"context"

	"github.com/LuanTenorio/learn-api/internal/auth/dto"
)

type AuthUseCase interface {
	Login(ctx context.Context, loginDto *dto.LoginDTO) (string, error)
}
