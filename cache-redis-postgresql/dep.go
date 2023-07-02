package main

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

type Dep struct {
	Db    *sql.DB
	Cache *redis.Client
}

func NewDep() *Dep {

	return &Dep{}
}
