package handler

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Login(c echo.Context) error
}
