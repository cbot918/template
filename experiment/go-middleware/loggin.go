package main

import (
	"database/sql"
	"fmt"
	"time"
)

type LogginService struct {
	next Service
}

func NewLogginService(next Service) Service {
	return &LogginService{
		next: next,
	}
}

func (l *LogginService) Ping() string {
	err := error.Error(sql.ErrNoRows)
	fmt.Printf("msg=%s err=%v took=%v\n", "hoooo", err, time.Since(time.Now()))
	return l.next.Ping()
}
