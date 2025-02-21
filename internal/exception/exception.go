package exception

import "github.com/labstack/echo/v4"

type Exception interface {
	AddTraceLog(info string) Exception
	Error() string
	HttpException(c echo.Context) error
	GetTrace() []string
}
