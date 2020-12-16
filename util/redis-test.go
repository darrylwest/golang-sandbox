package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

func main() {
	fmt.Println("redis tests")
	opts := new(redis.Options)
	opts.Addr = "localhost:6379"

	client := redis.NewClient(opts)

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("GET", key)
		client.Process(cmd)
		return cmd
	}

	key := "fb:10152879999728524:name"
	v, err := Get(client, key).Result()
    if err != nil {
        fmt.Printf("get key: %s, error: %s\n", key, err)
    } else {
        fmt.Printf("get key: %s, result: %s\n", key, v)
    }

	/// -- or

	pong, err := client.Ping().Result()
    if err != nil {
        fmt.Printf("ping error: %s\n", err)
    } else {
        fmt.Println(pong)
    }

	/// -- or

	r, err := client.Get(key).Result()
    if err != nil {
        fmt.Printf("err: %s\n", err)
    } else {
        fmt.Printf("result: %s\n", r)
    }
	/// -- or

	x, err := client.Get("bad-key").Result()
	if err == redis.Nil {
		fmt.Println("key 'bad-key' does not exist...")
	}

	if x == "" {
		fmt.Printf("result: (empty string), err: %s\n", err)
	} else {
		fmt.Printf("result: %s, err: %s\n", x, err)
	}
}
