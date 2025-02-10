package server

import (
	"fmt"

	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const ApiPrefix = "api"

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Validator = &CustomValidator{validator: validator.New()}
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK\n")
	})

	s.bootHandlers()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
