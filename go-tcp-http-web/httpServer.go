package main

import (
	"net"
	"net/http"
)

type HttpServer struct {
	Router http.ServeMux
	l      net.Listener
}

func NewHttpServer(l net.Listener) *HttpServer {
	qs := new(HttpServer)

	qs.Router.HandleFunc("/ping", qs.Ping)
	qs.Router.Handle("/", http.FileServer(http.Dir(".")))
	return qs
}

func (qs *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qs.Router.ServeHTTP(w, r)
}

func (qs *HttpServer) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
