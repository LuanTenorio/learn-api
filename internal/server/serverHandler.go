package server

import (
	authHandler "github.com/LuanTenorio/learn-api/internal/auth/handler"
	authUsecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
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
