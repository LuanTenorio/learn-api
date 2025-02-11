package util

import (
	"net/http"

	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/labstack/echo/v4"
)

func BindBody[T any](c echo.Context, dto T) exception.Exception {
	if err := c.Bind(dto); err != nil {
		return exception.New(exception.IncompatibleBody, http.StatusBadRequest, err.Error())
	} else if err := c.Validate(dto); err != nil {
		return exception.New(err.Error(), http.StatusBadRequest)
	}

	return nil
}
