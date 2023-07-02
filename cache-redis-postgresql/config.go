package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	DB_TYPE   string
	DB_URL    string
	CACHE_URL string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log("godotenv load failed")
		panic(err)
	}

	return &Config{
		PORT:      os.Getenv("PORT"),
		DB_TYPE:   os.Getenv("DB_TYPE"),
		DB_URL:    os.Getenv("DB_URL"),
		CACHE_URL: os.Getenv("CACHE_URL"),
	}
}
