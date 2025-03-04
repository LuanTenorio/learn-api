package handler

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/labstack/echo/v4"
)

// @Summary		List subjects
// @Description	List many subjects
// @Tags			Subject
// @Accept			json
// @Produce		json
// @Param			request	query		pagination.PaginationRequestDTO	true	"Optional data for pagination"
// @Success		200		{object}	pagination.PaginationResponseDTO
// @Failure		401		{object}	exception.ExceptionImpl "Unauthorized user"
// @Failure		500		{object}	exception.ExceptionImpl
// @Router			/subjects [get]
func (s *subjectHandlerImpl) List(c echo.Context) error {
	pgt, userId, err := pagination.GetPaginationAndUserId(c)

	if err != nil {
		return err
	}

	pgtResp, err := s.usecase.List(c.Request().Context(), pgt, userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pgtResp)
}
