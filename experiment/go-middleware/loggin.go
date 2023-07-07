package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("k", "g").Str("m", "r").Int("code", 300).Msg("auth service")
	// err := error.Error(sql.ErrNoRows)
	// fmt.Printf("msg=%s err=%v took=%v\n", "hoooo", err, time.Since(time.Now()))
	return l.next.Ping()
}
