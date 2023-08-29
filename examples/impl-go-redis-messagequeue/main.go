package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	for {
		// Pop a message from the queue (blocking)
		result, err := client.BRPop(ctx, 0, "myqueue").Result()
		if err != nil {
			fmt.Println("Error popping message:", err)
		} else {
			message := result[1]
			fmt.Println("Consumed:", message)
		}
	}
}
