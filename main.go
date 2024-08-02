package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/titosunu/wallet-go/infrastructure/api"
	"github.com/titosunu/wallet-go/infrastructure/component"
	"github.com/titosunu/wallet-go/infrastructure/config"
	"github.com/titosunu/wallet-go/infrastructure/middleware"
	"github.com/titosunu/wallet-go/infrastructure/repository"
	"github.com/titosunu/wallet-go/infrastructure/services"
)

func main() {
	config := config.Get()
	dbConnection := component.GetDatabaseConnection(config)
	cacheConnection := component.GetCacheConnection()
	userRepository := repository.NewUser(dbConnection)
	userService := services.NewUser(userRepository, cacheConnection)
	authMid := middleware.Authenticate(userService)

	app := fiber.New()

	api.NewAuth(app, userService, authMid)
	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}