package main

// redistogo.com
// darryl.west@raincitysoftware.com
// 5MB

// don't know where it's hosted

import (
	"fmt"
	"github.com/darrylwest/naledi-key-service/src/keyservice"
	"gopkg.in/redis.v3"
	"os"
	"path"
)

func main() {
	fmt.Println("redistogo tests")

	filename := path.Join(os.Getenv("HOME"), ".keyservice", "config.json")
	config, err := keyservice.ReadConfig(filename)

	client := redis.NewClient(config.GetSecondaryRedisOptions())

	Get := func(client *redis.Client, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd("GET", key)
		client.Process(cmd)
		return cmd
	}

	/// -- or

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("connect error: ", err)
	} else {
		fmt.Println(pong)

		key := "Project:12345"
		v, err := Get(client, key).Result()
		fmt.Printf("result: %s, err: %s\n", v, err)

		/// -- or

		r, err := client.Get(key).Result()
		fmt.Printf("result: %s, err: %s\n", r, err)
		/// -- or

		x, err := client.Get("bad-key").Result()
		if err == redis.Nil {
			fmt.Println("key does not exist, redis.Nil...")
		} else {
			fmt.Printf("result: (empty string), err: %s\n", err)
			fmt.Println(x)
		}

	}
}
