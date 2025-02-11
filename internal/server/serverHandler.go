package server

import (
	"fmt"
	"net/http"

	authHandler "github.com/LuanTenorio/learn-api/internal/auth/handler"
	authUsecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
	"github.com/LuanTenorio/learn-api/internal/exception"
	userHandler "github.com/LuanTenorio/learn-api/internal/user/handler"
	userRepository "github.com/LuanTenorio/learn-api/internal/user/repository"
	userUseCase "github.com/LuanTenorio/learn-api/internal/user/useCase"
	"github.com/labstack/echo/v4"
)

func (s *echoServer) bootHandlers() {
	s.app.GET(ApiPrefix+"/ping", func(c echo.Context) error {
		return c.JSON(200, "Pong")
	})

	userRepo := bootUserHandler(s)
	bootAuthHandler(s, userRepo)
}

func bootUserHandler(s *echoServer) userRepository.UserRepository {
	userRepo := userRepository.NewUserPGRepository(s.db)
	userUC := userUseCase.NewUserUseCaseImpl(userRepo)
	userHand := userHandler.NewUserUseCaseImpl(userUC)

	userRoutes := s.app.Group(ApiPrefix + "/users")

	userRoutes.POST("", userHand.CreateUser)

	return userRepo
}

func bootAuthHandler(s *echoServer, userRepo userRepository.UserRepository) {
	authUC := authUsecase.NewAuthUseCaseImpl(userRepo)
	authHand := authHandler.NewAuthHandlerImpl(authUC)

	authRoutes := s.app.Group(ApiPrefix + "/auth")

	authRoutes.POST("/login", authHand.Login)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	respErro, ok := err.(*exception.ExceptionImpl)

	if !ok {
		respErro = exception.New("Internal server error", http.StatusInternalServerError, err.Error())
	}

	fmt.Println(err.Error())
	fmt.Println(respErro.Trace)

	respErro.HttpException(c)
}
