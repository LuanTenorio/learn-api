package server

import (
	userHandler "github.com/LuanTenorio/learn-api/internal/user/handler"
	userRepository "github.com/LuanTenorio/learn-api/internal/user/repository"
	userUseCase "github.com/LuanTenorio/learn-api/internal/user/useCase"
)

func (s *echoServer) bootHandlers() {
	bootUserHandler(s)
}

func bootUserHandler(s *echoServer) {
	userRepo := userRepository.NewUserPGRepository(s.db)
	userUc := userUseCase.NewUserUseCaseImpl(userRepo)
	userHand := userHandler.NewUserUseCaseImpl(userUc)

	const ApiRouterBase = ApiPrefix + "/users"
	userRoutes := s.app.Group(ApiRouterBase)

	userRoutes.POST("", userHand.CreateUser)
}
