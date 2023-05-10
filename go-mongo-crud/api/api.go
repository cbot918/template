package api

import (
	"github.com/cbot918/go-mongo-crud/api/controller"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct{
	Server *fiber.App
}

func InitApi(mclient *mongo.Client) *fiber.App{
	api := fiber.New()

	// regist routers
	c := controller.New(mclient)
	api.Get("/ping", c.Ping)
	api.Get("/use/:dbname/:coname", c.UseCollection)
	
	api.Post("/insert", c.Insert)
	// api.Post("/insertmany")

	return api
}
