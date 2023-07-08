package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

var (
	lg = fmt.Println
	lf = fmt.Printf
)

type Config struct {
	HOST     string
	PSQL_URL string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	return &Config{
		HOST:     os.Getenv("HOST"),
		PSQL_URL: os.Getenv("PSQL_URL"),
	}, nil
}

func GetUserFromEmail(email string) (string, error) {
	reg, err := regexp.Compile(".*@")
	if err != nil {
		return "", err
	}
	return strings.Trim(reg.FindString(email), "@"), nil
}
