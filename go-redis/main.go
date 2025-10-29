package main

import "github.com/go-redis/redis"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client.Get("name")
}
