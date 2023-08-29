package main

import (
	"log"

	i "officialweb/internal"

	"github.com/gofiber/fiber/v2"
)

const port = ":8081"

func main() {

	cfg, err := i.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := i.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	f := fiber.New()
	api := i.NewAPI(cfg, f, db)

	log.Fatal(api.App.Listen(port))
}
