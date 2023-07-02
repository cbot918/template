package main

import (
	"net"
	"net/http"
	"os"

	"github.com/cbot918/template/go-tcp-http-web/pkg/wsy"
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

	// // run ws server
	wsyServer := wsy.New(listener)
	wsyServer.Run()

	// run http server
	// qhs := NewQHttpServer(listener)
	// if err := qhs.Serve(listener); err != nil {
	// 	log("http serve failed")
	// 	panic(err)
	// }

	// // tcp server
	// RunTcpServer(listener)
}

func NewQHttpServer(l net.Listener) *http.Server {
	s := &http.Server{
		Handler: NewHttpServer(l),
	}
	return s
}
