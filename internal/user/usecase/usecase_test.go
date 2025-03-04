package usecase_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/user/dto"
	"github.com/LuanTenorio/learn-api/internal/user/entity"
	usecase "github.com/LuanTenorio/learn-api/internal/user/usecase"
	"github.com/LuanTenorio/learn-api/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userUC := usecase.NewUserUseCaseImpl(mockRepo)

	ctx := context.Background()
	userDto := &dto.CreateUserDTO{Name: "Thierry Luan", Email: "thierry@example.com", Password: "password123"}
	expectedUser := &entity.User{Id: 1, Name: "Thierry Luan", Email: "thierry@example.com"}

	mockRepo.On("CreateUser", ctx, mock.Anything).Return(expectedUser, nil)
	user, err := userUC.CreateUser(ctx, userDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_ErrorOnCreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userUC := usecase.NewUserUseCaseImpl(mockRepo)

	ctx := context.Background()
	userDto := &dto.CreateUserDTO{Name: "Thierry Luan", Email: "thierry@example.com", Password: "password123"}
	conflictUserError := exception.New("There is already a user with this email", http.StatusConflict)
	expectedError := &conflictUserError

	mockRepo.On("CreateUser", ctx, mock.Anything).Return(nil, conflictUserError)
	user, err := userUC.CreateUser(ctx, userDto)

	assert.Error(t, err)
	assert.Nil(t, user)
	exception.CheckExceptionForTest(t, err, **expectedError)

	mockRepo.AssertExpectations(t)
}
