package util

import (
	"errors"

	"github.com/LuanTenorio/learn-api/internal/requestError"
	"github.com/labstack/echo/v4"
)

func BindBody[T any](c echo.Context, dto T) error {
	if err := c.Bind(dto); err != nil {
		return errors.New(requestError.IncompatibleBody)
	} else if err := c.Validate(dto); err != nil {
		return err
	}

	return nil
}
