package handler

import "github.com/labstack/echo/v4"

type SubjectHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
}
