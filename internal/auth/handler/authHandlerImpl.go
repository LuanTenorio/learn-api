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

// @Summary		Login
// @Description	Authenticate a user and returns a JWT token
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		dto.LoginDTO	true	"Login data"
// @Success		200		{object}	dto.LoginResponseDTO
// @Failure		404		{object}	exception.ExceptionImpl	"No user with this email"
// @Failure		401		{object}	exception.ExceptionImpl	"Wrong password"
// @Failure		500		{object}	exception.ExceptionImpl
// @Router			/auth/login [post]
func (h *authHandlerImpl) Login(c echo.Context) error {
	loginDto := new(dto.LoginDTO)

	if err := util.BindBody(c, loginDto); err != nil {
		return err
	}

	token, err := h.AuthUC.Login(c.Request().Context(), loginDto)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &dto.LoginResponseDTO{Token: token})
}
