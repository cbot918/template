package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PSQL_URL string
	HOST     string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		PSQL_URL: os.Getenv("PSQL_URL"),
		HOST:     os.Getenv("HOST"),
	}, nil
}

func GetNameFromEmail(email string) string {
	reg, err := regexp.Compile(".*@")
	if err != nil {
		log.Fatal(err)
	}
	res := reg.Find([]byte(email))
	return strings.TrimRight(string(res), "@")
}
