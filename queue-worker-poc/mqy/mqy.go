package main

import "net/http"

type Mqy struct {
	router http.ServeMux
}

func NewMqy() (mqy *Mqy) {

	mqy = new(Mqy)

	mqy.router.HandleFunc("/", Ping)
	mqy.router.HandleFunc("/mq", MQ)

	return
}

func (m *Mqy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

func MQ(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mq\n"))
}
