package main

import (
	"github.com/gofiber/fiber/v2"
	conf "koriebruh/restful/api/config"
	"koriebruh/restful/api/handler"
	"koriebruh/restful/api/middleware"
	"koriebruh/restful/api/repository"
	"koriebruh/restful/api/service"
	"koriebruh/restful/api/utils"
	"time"
)

func main() {

	dbConnection := conf.InitDB()
	cacheConnection := utils.GetCacheConnection()
	repositoryAuth := repository.NewAuthRepository(dbConnection)
	authService := service.NewAuthService(repositoryAuth, cacheConnection)

	UserMid := middleware.Authentication(authService)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})
	handler.NewAuth(app, authService, UserMid)

	app.Listen(":3000")
}

func HelloHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World")
}
