package server

import (
	authHandler "github.com/LuanTenorio/learn-api/internal/auth/handler"
	authUsecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
	userHandler "github.com/LuanTenorio/learn-api/internal/user/handler"
	userRepository "github.com/LuanTenorio/learn-api/internal/user/repository"
	userUseCase "github.com/LuanTenorio/learn-api/internal/user/useCase"
)

func (s *echoServer) bootHandlers() {
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
