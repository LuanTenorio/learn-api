package handler

import "github.com/labstack/echo/v4"

type UserHandler interface {
	CreateUser(c echo.Context) error
}
