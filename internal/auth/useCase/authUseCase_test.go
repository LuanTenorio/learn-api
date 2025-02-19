package usecase_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/LuanTenorio/learn-api/internal/auth/dto"
	usecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
	"github.com/LuanTenorio/learn-api/internal/exception"
	userDto "github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Sucess(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	userUC := usecase.NewAuthUseCaseImpl(mockUserRepo)

	ctx := context.Background()
	loginDto := &dto.LoginDTO{Email: "thierry@example.com", Password: "@Pwd123456"}
	expectedUser := &userDto.UserWithPwdDTO{Id: 1, Name: "Thierry Luan", Email: "thierry@example.com", Password: "$2a$10$Hrrek2suWGiVTDvcH6pHHuJl5QEyXSYxgCYchnwwLXaGZk.M9yJXu"}

	mockUserRepo.On("FindUserAndPwdByEmail", ctx, loginDto.Email).Return(expectedUser, nil)
	token, err := userUC.Login(ctx, loginDto)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockUserRepo.AssertExpectations(t)
}

func TestLogin_ErrorPwd(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	userUC := usecase.NewAuthUseCaseImpl(mockUserRepo)

	ctx := context.Background()
	loginDto := &dto.LoginDTO{Email: "thierry@example.com", Password: "@Senha1poads"}
	expectedUser := &userDto.UserWithPwdDTO{Id: 1, Name: "Thierry Luan", Email: "thierry@example.com", Password: "$2a$10$Hrrek2suWGiVTDvcH6pHHuJl5QEyXSYxgCYchnwwLXaGZk.M9yJXu"}
	expectedException := exception.New("wrong password", http.StatusUnauthorized)

	mockUserRepo.On("FindUserAndPwdByEmail", ctx, loginDto.Email).Return(expectedUser, nil)
	token, err := userUC.Login(ctx, loginDto)

	assert.Empty(t, token)
	assert.Error(t, err)
	exception.CheckExceptionForTest(t, err, *expectedException)

	mockUserRepo.AssertExpectations(t)
}

func TestLogin_ErrorOnNotFoundUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	userUC := usecase.NewAuthUseCaseImpl(mockUserRepo)

	ctx := context.Background()
	loginDto := &dto.LoginDTO{Email: "thierry@example.com", Password: "@Pwd123456"}
	repoException := exception.New("There is no user with this email", http.StatusNotFound)
	expectedException := &repoException

	mockUserRepo.On("FindUserAndPwdByEmail", ctx, loginDto.Email).Return(nil, repoException)
	token, err := userUC.Login(ctx, loginDto)

	assert.Empty(t, token)
	assert.Error(t, err)
	exception.CheckExceptionForTest(t, err, **expectedException)

	mockUserRepo.AssertExpectations(t)
}
