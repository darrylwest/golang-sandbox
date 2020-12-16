package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

func main() {
	fmt.Println("redis tests")
	opts := new(redis.Options)
	opts.Addr = "127.0.0.1:6380"

	client := redis.NewClient(opts)

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("GET", key)
		client.Process(cmd)
		return cmd
	}

	key := "mykey"
	v, err := Get(client, key).Result()
	fmt.Printf("result: %s, err: %s\n", v, err)

	/// -- or

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	/// -- or

	r, err := client.Get(key).Result()
	fmt.Printf("result: %s, err: %s\n", r, err)
	/// -- or

	x, err := client.Get("bad-key").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist, redis.Nil...")
	}

	if x == "" {
		fmt.Printf("result: (empty string), err: %s\n", err)
	} else {
		fmt.Printf("result: %s, err: %s\n", x, err)
	}
}
