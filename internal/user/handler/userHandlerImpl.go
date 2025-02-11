package handler

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
	usecase "github.com/LuanTenorio/learn-api/internal/user/useCase"
	"github.com/LuanTenorio/learn-api/internal/util"
	"github.com/labstack/echo/v4"
)

const (
	errorOnCreateUser = "Error creating user"
)

type userHandlerImpl struct {
	UserUC usecase.UserUseCase
}

func NewUserUseCaseImpl(uc usecase.UserUseCase) UserHandler {
	return &userHandlerImpl{UserUC: uc}
}

func (h *userHandlerImpl) CreateUser(c echo.Context) error {
	userDto := new(dto.CreateUserDTO)

	if err := util.BindBody(c, userDto); err != nil {
		return err
	}

	user, err := h.UserUC.CreateUser(c.Request().Context(), userDto)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
