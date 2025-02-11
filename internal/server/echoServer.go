package server

import (
	"encoding/json"
	"fmt"

	"github.com/LuanTenorio/learn-api/internal/auth"
	"github.com/LuanTenorio/learn-api/internal/config"
	"github.com/LuanTenorio/learn-api/internal/database"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const ApiPrefix = "api"

type echoServer struct {
	app    *echo.Echo
	db     database.Database
	conf   *config.Config
	jwtMid echo.MiddlewareFunc
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
		app:    echoApp,
		db:     db,
		conf:   conf,
		jwtMid: getJWTMiddleware(),
	}
}

func (s *echoServer) Start() {
	s.app.Validator = &CustomValidator{validator: validator.New()}
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.HTTPErrorHandler = customHTTPErrorHandler

	s.bootHandlers()
	showRoutes(s)

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func showRoutes(e *echoServer) {
	data, err := json.MarshalIndent(e.app.Routes(), "", "  ")
	if err != nil {
		return
	}

	fmt.Println(string(data))
}

func getJWTMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(auth.SecretKey),
	}

	return echojwt.WithConfig(config)
}
