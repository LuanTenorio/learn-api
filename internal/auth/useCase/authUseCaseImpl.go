package usecase

import (
	"github.com/LuanTenorio/learn-api/internal/user/repository"
)

type authUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewAuthUseCaseImpl(ur repository.UserRepository) AuthUseCase {
	return &authUseCaseImpl{userRepository: ur}
}
