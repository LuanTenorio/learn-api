package handler

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/user/dto"
	usecase "github.com/LuanTenorio/learn-api/internal/user/useCase"
	"github.com/LuanTenorio/learn-api/internal/util"
	"github.com/labstack/echo/v4"
)

type userHandlerImpl struct {
	UserUC usecase.UserUseCase
}

func NewUserUseCaseImpl(uc usecase.UserUseCase) UserHandler {
	return &userHandlerImpl{UserUC: uc}
}

// @Summary		New User
// @Description	Creates a new user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			request	body		dto.CreateUserDTO	true	"Data required for the user's acriation"
// @Success		200		{object}	entity.User
// @Failure		409		{object}	exception.ExceptionImpl	"There is already a user with this email"
// @Failure		500		{object}	exception.ExceptionImpl
// @Router			/users [post]
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
