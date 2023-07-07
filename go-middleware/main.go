package main

import (
	"fmt"
	"log"
	"net/http"
)

var lg = fmt.Println

func main() {

	svc := NewPingService()
	svc = NewLogginService(svc)
	svc = NewAuthService(svc)
	svc.Ping()

	router := http.NewServeMux()

	router.HandleFunc("/ping/kk", func(w http.ResponseWriter, r *http.Request) {
		lg("kk")
	})

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		lg("pong")
	})

	handler := AuthMiddleware(router)

	log.Fatal(http.ListenAndServe(":12345", handler))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lg("authorize process")
		next.ServeHTTP(w, r)
	})
}
