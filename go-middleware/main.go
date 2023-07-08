package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var lg = fmt.Println

func main() {

	svc := NewPingService()
	svc = NewLogginService(svc)
	svc = NewAuthService(svc)
	svc.Ping()

	// router := http.NewServeMux()

	r := mux.NewRouter()

	r.HandleFunc("/ping/kk", func(w http.ResponseWriter, r *http.Request) {
		lg("kk")
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		lg("pong")
	})

	handler := AuthMiddleware(r)

	log.Fatal(http.ListenAndServe(":12345", handler))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lg("authorize process")
		next.ServeHTTP(w, r)
	})
}
