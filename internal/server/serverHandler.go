package server

import (
	"net/http"

	authHandler "github.com/LuanTenorio/learn-api/internal/auth/handler"
	"github.com/LuanTenorio/learn-api/internal/auth/middleware"
	authUsecase "github.com/LuanTenorio/learn-api/internal/auth/useCase"
	"github.com/LuanTenorio/learn-api/internal/exception"
	"github.com/LuanTenorio/learn-api/internal/logger"
	subjectHandler "github.com/LuanTenorio/learn-api/internal/subject/handler"
	subjectRepository "github.com/LuanTenorio/learn-api/internal/subject/repository"
	subjectUseCase "github.com/LuanTenorio/learn-api/internal/subject/usecase"
	userHandler "github.com/LuanTenorio/learn-api/internal/user/handler"
	userRepository "github.com/LuanTenorio/learn-api/internal/user/repository"
	userUseCase "github.com/LuanTenorio/learn-api/internal/user/useCase"
	"github.com/labstack/echo/v4"
)

var serverHandlerLogger = logger.New("server", "handler")

func (s *echoServer) bootHandlers() {
	s.app.GET(ApiPrefix+"/ping", func(c echo.Context) error {
		return c.JSON(200, "Pong")
	})

	userRepo := bootUserHandler(s)
	bootAuthHandler(s, userRepo)
	bootSubjectHandler(s)
}

func bootUserHandler(s *echoServer) userRepository.UserRepository {
	userRepo := userRepository.NewUserPGRepository(s.db)
	userUC := userUseCase.NewUserUseCaseImpl(userRepo)
	userHand := userHandler.NewUserHandlerImpl(userUC)

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

func bootSubjectHandler(s *echoServer) {
	subjectRepo := subjectRepository.New(s.db)
	subjectUC := subjectUseCase.New(subjectRepo)
	subjectHand := subjectHandler.New(subjectUC)

	subjectRoutes := s.app.Group(ApiPrefix+"/subjects", middleware.AuthMiddleware)

	subjectRoutes.POST("", subjectHand.Create)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	respErro, ok := err.(*exception.ExceptionImpl)

	if !ok {
		respErro = exception.New("Internal server error", http.StatusInternalServerError, err.Error())
	}

	serverHandlerLogger.Error(err.Error())
	for _, log := range respErro.GetTrace() {
		serverHandlerLogger.Debug(log)
	}

	respErro.HttpException(c)
}
