package server

import (
	"github.com/cbot918/grpost/server/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRouting(server *fiber.App) {
	// health check
	server.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"code": 200, "message": "Hello, World"})
	})

	c := controller.NewController()
	server.Post("/signin", c.Auth.Signin)
}
