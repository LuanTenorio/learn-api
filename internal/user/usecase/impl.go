package usecase

import (
	"github.com/LuanTenorio/learn-api/internal/user/repository"
)

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(ur repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{userRepository: ur}
}
