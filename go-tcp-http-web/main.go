package main

import (
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log("load env failed")
		panic(err)
	}

	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log("net listen failed")
		os.Exit(1)
	}
	log("listening: ", port)
	defer listener.Close()

	// run http server
	qhs := NewQHttpServer()
	if err := qhs.Serve(listener); err != nil {
		log("http serve failed")
		panic(err)
	}

	// // tcp server
	// RunTcpServer(listener)

}

func NewQHttpServer() *http.Server {
	s := &http.Server{
		Handler: NewHttpServer(),
	}
	return s
}
