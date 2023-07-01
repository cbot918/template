package main

import "net/http"

type HttpServer struct {
	Router http.ServeMux
}

func NewHttpServer() *HttpServer {
	qs := new(HttpServer)

	qs.Router.HandleFunc("/ping", pingController)
	qs.Router.Handle("/", http.FileServer(http.Dir(".")))

	return qs
}

func (q *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q.Router.ServeHTTP(w, r)
}

func pingController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
