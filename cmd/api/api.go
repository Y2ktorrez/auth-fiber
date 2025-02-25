package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mutinho/src"
)

func SetupApi(app *fiber.App, c *src.Container) {
	api := app.Group("/api/v1")

	handlers := []func(fiber.Router){
		//UserHandler
		c.UserHandler.RegisterRoutes,
	}

	for _, register := range handlers {
		register(api)
	}

}
