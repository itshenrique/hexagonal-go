package main

import (
	"fmt"
	"nito/api/internals/config"
	"nito/api/internals/core/services"
	"nito/api/internals/handlers"
	"nito/api/internals/middlewares"
	"nito/api/internals/repositories"
	"nito/api/internals/server"
)

func main() {
	config := config.LoadConfig(".")

	// Database config
	mariaDbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
		config.MySQLDatabase,
	)

	redisAddress := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	// Middlewares
	loggerMiddleware := middlewares.NewLoggerMiddleware()
	authRequiredMiddleware := middlewares.NewAuthRequiredMiddleware(
		services.NewSessionService(
			repositories.NewSessionRepository(redisAddress),
		),
	)

	// Handlers
	accountHandler := handlers.NewAccountHandler(
		services.NewAccountService(
			repositories.NewAccountRepository(mariaDbConnectionString),
		),
		services.NewSessionService(
			repositories.NewSessionRepository(redisAddress),
		),
	)

	characterHandler := handlers.NewCharacterHandler(
		services.NewCharacterService(
			repositories.NewCharacterRepository(mariaDbConnectionString),
		),
	)

	httpServer := server.NewServer(
		loggerMiddleware,
		authRequiredMiddleware,
		accountHandler,
		characterHandler,
	)

	httpServer.Initialize()
}
