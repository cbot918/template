package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("a", "a").Int("status_code", 200).Msg("k")

	// log.Str("protocol", "grpc").
	// 	Str("method", info.FullMethod).
	// 	Int("status_code", int(statusCode)).
	// 	Str("status_text", statusCode.String()).
	// 	Dur("duration", duration).
	// 	Msg("received a gRPC request")
	// log.Info().Str("me", "gg").
	// log.Print("hihi")
}
