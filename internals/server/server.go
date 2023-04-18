package server

import (
	"nito/api/internals/core/ports"
	"nito/api/internals/middlewares"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Initialize()
}

type Server struct {
	// Middlewares
	loggerMiddleware       middlewares.ILoggerMiddleware
	authRequiredMiddleware middlewares.IAuthRequiredMiddleware

	// Handlers
	accountHandler   ports.IAccountHandler
	characterHandler ports.ICharacterHandler
}

func NewServer(
	loggerMiddleware middlewares.ILoggerMiddleware,
	authRequiredMiddleware middlewares.IAuthRequiredMiddleware,
	accountHandler ports.IAccountHandler,
	characterHandler ports.ICharacterHandler,
) IServer {
	return &Server{
		loggerMiddleware:       loggerMiddleware,
		authRequiredMiddleware: authRequiredMiddleware,
		accountHandler:         accountHandler,
		characterHandler:       characterHandler,
	}
}

func (s *Server) Initialize() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// TODO - melhorar logger
	// router.Use(s.loggerMiddleware.LoggerMiddleware())
	// accountHandler
	router.POST("/api/auth/login", s.accountHandler.Login)
	router.POST("/api/auth/logout", s.authRequiredMiddleware.Initialize(), s.accountHandler.Logout)
	router.GET("/api/account", s.authRequiredMiddleware.Initialize(), s.accountHandler.GetAllAccounts)
	router.POST("/api/account", s.accountHandler.CreateAccount)
	router.GET("/api/account/:id", s.authRequiredMiddleware.Initialize(), s.accountHandler.FindById)

	// characterHandler
	router.GET("/api/character", s.authRequiredMiddleware.Initialize(), s.characterHandler.GetAllCharacters)
	router.POST("/api/character", s.authRequiredMiddleware.Initialize(), s.characterHandler.CreateCharacter)
	router.GET("/api/character/:id", s.authRequiredMiddleware.Initialize(), s.characterHandler.FindById)

	router.Run("localhost:3001")
}
