package rest

import "github.com/gofiber/fiber/v2"

type Rest struct{}

func New() *Rest {
	return &Rest{}
}

func (r *Rest) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("Pong!!")
}
