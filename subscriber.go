package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	channel = "EXAMPLE"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}

	pubsub := client.Subscribe(channel)
	_, err = pubsub.Receive()
	if err != nil {
		fmt.Println(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println("Received: ", msg.Payload)
		if msg.Payload == "-1" {
			break
		}
	}

	_ = pubsub.Close()
	fmt.Println("Pubsub closed. Channel closed.")
}
