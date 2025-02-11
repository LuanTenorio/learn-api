package handler

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/auth/dto"
	usecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
	"github.com/LuanTenorio/learn-api/internal/util"
	"github.com/labstack/echo/v4"
)

type authHandlerImpl struct {
	AuthUC usecase.AuthUseCase
}

func NewAuthHandlerImpl(uc usecase.AuthUseCase) AuthHandler {
	return &authHandlerImpl{AuthUC: uc}
}

func (h *authHandlerImpl) Login(c echo.Context) error {
	loginDto := new(dto.LoginDTO)

	if err := util.BindBody(c, loginDto); err != nil {
		return err
	}

	token, err := h.AuthUC.Login(c.Request().Context(), loginDto)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
