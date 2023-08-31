package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_TYPE string
	DB_URL  string
	PORT    string
}

func NewConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		DB_TYPE: os.Getenv("DB_TYPE"),
		DB_URL:  os.Getenv("DB_URL"),
		PORT:    os.Getenv("PORT"),
	}, nil
}

// type Config struct {
// 	DB struct {
// 		Username string
// 		Password string
// 		Host     string
// 		Port     string
// 		DB_NAME  string
// 	}
// }
