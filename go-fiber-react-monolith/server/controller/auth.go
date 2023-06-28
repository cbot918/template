package controller

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/cbot918/liby/jwty"
	"github.com/gofiber/fiber/v2"
)

type Auth struct{}

func NewAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Signin(c *fiber.Ctx) error {

	user := struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&user); err != nil {
		fmt.Println("err via bodyparser")
		panic(err)
	}

	// jwt process
	token, err := jwty.New().FastJwt(int(user.Id), user.Email)
	if err != nil {
		fmt.Println("err call FastJwt")
		panic(err)
	}

	res := struct {
		Token string `json:"token"`
		User  string `json:"user"`
	}{}

	res.Token = token
	res.User = strings.Trim(regexp.MustCompile(".*@").FindString(user.Email), "@")

	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Println("err via json Marshal")
	}

	return c.SendString(string(resp))

}
