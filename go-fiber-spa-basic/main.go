package main

import (
	"net/http"

	"github.com/cbot918/go-fiber-serve-react-spa/rest"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Static("/", "./ui/dist")

	// router
	rest := rest.New()
	app.Get("/ping", rest.Ping)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})

	err := app.Listen(":3001")
	if err != nil {
		panic(err)
	}
}
