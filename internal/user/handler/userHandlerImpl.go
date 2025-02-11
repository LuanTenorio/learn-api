package handler

import (
	"log"
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/requestError"
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
		return c.JSON(http.StatusBadRequest, requestError.New(err.Error()))
	}

	user, err := h.UserUC.CreateUser(c.Request().Context(), userDto)

	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, requestError.New(errorOnCreateUser))
	}

	return c.JSON(http.StatusCreated, user)
}
