package main

import (
	"github.com/gofiber/fiber/v2"
	conf "koriebruh/restful/api/config"
	"time"
)

func main() {

	dbConnnection := conf.InitDB()

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	app.Get("/", HelloHandler)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}

func HelloHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World")
}
