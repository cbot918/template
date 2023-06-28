package main

import (
	"github.com/cbot918/grpost/server"
)

const (
	port       = ":5010"
	staticPath = "ui/build"
)

type Config struct {
	PORT       string
	StaticPath string
}

func NewConfig() *Config {
	return &Config{
		PORT: port,
	}
}

func main() {

	cfg := NewConfig()

	grpost := server.New(cfg.StaticPath)

	grpost.Server.Listen(cfg.PORT)
}
