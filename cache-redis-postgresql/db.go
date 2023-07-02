package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDb(cfg *Config) (db *sql.DB) {
	db, err := sql.Open(cfg.DB_TYPE, cfg.DB_URL)
	if err != nil {
		log("sql.Open failed")
		panic(err)
	}
	return
}
