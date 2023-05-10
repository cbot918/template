package controller

import (
	"fmt"
	"strings"

	"encoding/json"
	"regexp"

	"github.com/cbot918/liby/jwty"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {}

func New() *Controller{
	return &Controller{}
}

func (ctlr *Controller) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func (ctlr *Controller) Auth(c *fiber.Ctx) error {
	// fmt.Println("in auth")
	// return c.SendString("got")
	user := struct{
		Id int `json:"id"`
		Email string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&user); err != nil { panic(err) }

	
	token, err:= jwty.New().FastJwt(int(user.Id), user.Email); if err != nil {panic(err)}


	res := struct {
		Token string `json:"token"`
		User string `json:"user"`
	}{}
	res.Token = token
	res.User = strings.Trim(regexp.MustCompile(".*@").FindString(user.Email),"@")
	fmt.Printf("user: %s", res.User)
	resp,err := json.Marshal(res); if err != nil { panic(err) }

	return c.SendString(string(resp))
}