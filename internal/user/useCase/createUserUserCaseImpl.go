package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
	"github.com/LuanTenorio/learn-api/internal/user/repository"
)

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(ur repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{userRepository: ur}
}

func (uc *userUseCaseImpl) CreateUser(ctx context.Context, userDto *dto.CreateUserDTO) (*entity.User, error) {
	hashedPassword, err := genHash(userDto.Password)

	if err != nil {
		return nil, err
	}

	userDto.Password = hashedPassword

	id, err := uc.userRepository.CreateUser(ctx, userDto)

	if err != nil {
		return nil, err
	}

	user := entity.NewUserByCreateDto(userDto, id)

	return user, nil
}

func genHash(pwd string) (string, error) {
	bPwd := []byte(pwd)

	hashedPassword, err := bcrypt.GenerateFromPassword(bPwd, bcrypt.DefaultCost)

	return string(hashedPassword), err
}
