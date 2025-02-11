package usecase

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
)

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

	return string(hashedPassword), exception.New("Internal Error on generate hash", http.StatusInternalServerError, err.Error())
}
