package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mutinho/cmd/api"
	"github.com/mutinho/config"
	"github.com/mutinho/src"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, FiberApi!")
	})

	c := src.SetupContainer()
	api.SetupApi(app, c)

	log.Println("Server running on port", config.ServerPort)
	app.Listen(":" + config.ServerPort)

}
