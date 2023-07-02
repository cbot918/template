package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var log = fmt.Println

func getPort() (port string) {
	if err := godotenv.Load(); err != nil {
		log("godotenv load failed")
		panic(err)
	}

	return os.Getenv("PORT")
}

func makeListener(port string) (lis net.Listener) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log("net listen failed")
		panic(err)
	}
	return
}

func runServer(mqy *Mqy, lis net.Listener, port string) {
	server := http.Server{
		Handler: mqy,
	}
	log("mqy serving on: ", port)
	if err := server.Serve(lis); err != nil {
		log("server serve failed")
		panic(err)
	}
}
