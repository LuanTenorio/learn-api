package handler

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/pagination"
	"github.com/labstack/echo/v4"
)

// @Summary		New Subject
// @Description	Creates a new subject
// @Tags			Subject
// @Accept			json
// @Produce		json
// @Param			request	body		dto.CreateSubjectDTO	true	"Data required for the subject's acriation"
// @Success		201		{object}	entity.Subject
// @Failure		409		{object}	exception.ExceptionImpl	"There is already a subject with this name"
// @Failure		400		{object}	exception.ExceptionImpl "Incompatible body"
// @Failure		401		{object}	exception.ExceptionImpl "Unauthorized user"
// @Failure		500		{object}	exception.ExceptionImpl
// @Router			/subjects [post]
func (s *subjectHandlerImpl) List(c echo.Context) error {
	pgt, userId, err := pagination.GetPaginationAndUserId(c)

	if err != nil {
		return err
	}

	pgtResp, err := s.usecase.List(c.Request().Context(), pgt, userId)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, pgtResp)
}
