package main

import (
	"github.com/cbot918/jwt-persistant/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	PORT = ":3500"
)

func main() {
	api := fiber.New()

	api.Use(cors.New())

	api.Static("/", "./ui/dist")

	c := controller.New()
	api.Get("/ping", c.Ping)
	api.Post("/auth", c.Auth)

	api.Listen(PORT)
}
