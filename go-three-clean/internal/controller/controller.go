package controller

import (
	"sample/internal/model"
	"sample/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Controller struct {
	S *service.Service
}

func NewController(db *sqlx.DB) *Controller {

	return &Controller{
		S: service.NewService(db),
	}
}

func (ctr *Controller) Singup(c *fiber.Ctx) error {
	// handle request
	req := &model.SignupRequest{}
	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := ctr.S.SignupService(req)
	if err != nil {
		return err
	}

	// handle response
	res := &model.SignupResponse{
		Email: req.Email,
		Name:  req.Name,
	}
	return c.JSON(res)
}

func (ctr *Controller) Insert(c *fiber.Ctx) error {
	return ctr.S.InsertService()
}
