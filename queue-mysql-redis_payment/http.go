package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

const PORT = ":8855"

func main() {

	http.HandleFunc("/payments", paymentsHandler)

	fmt.Println("listening", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}

func paymentsHandler(w http.ResponseWriter, req *http.Request) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	buf := new(bytes.Buffer)

	// Include a Validation logic here to sanitize the req.Body when working in a production environment
	buf.ReadFrom(req.Body)

	paymentDetails := buf.String()

	err := redisClient.RPush("payments", paymentDetails).Err()

	if err != nil {
		fmt.Fprintf(w, err.Error()+"\r\n")
	} else {
		fmt.Fprintf(w, "Payment details accepted successfully\r\n")
	}
}
