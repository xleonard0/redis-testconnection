package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {

	// Client Init
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	})

	// Test connection

	pong, err := redisClient.Ping().Result()

	if err != nil {
		fmt.Println("Could not connect to Redis server")
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	// Return pong msg if connection to redis server is successful

	fmt.Println(pong, "-", "The connection to the redis server was successful!")

}
