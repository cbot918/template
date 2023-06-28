package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	Server *fiber.App
}

func New(uiPath string) *App {

	app := new(App)

	app.Server = fiber.New()
	app.Server.Use(cors.New())     // setup use cors
	app.Server.Static("/", uiPath) // serve spa
	SetupRouting(app.Server)       // setup routing

	return app
}
