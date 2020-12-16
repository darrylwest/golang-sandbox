package main

// redislabs.com
// darryl.west@raincitysoftware.com
//

// redis-labs:
//  hosted on AWS, Azure, GCE and IBM
//  free account is 30MB
//  $7/mon for 100MB with replica, failover, backups, 4 dbs
//  $18/mon for 250MB with replica, failover, backups, 8 dbs

import (
	"fmt"
	"github.com/darrylwest/naledi-key-service/src/keyservice"
	"gopkg.in/redis.v3"
	"os"
	"path"
)

func main() {
	fmt.Println("redis-labs tests")

	filename := path.Join(os.Getenv("HOME"), ".keyservice", "config.json")
	config, err := keyservice.ReadConfig(filename)

	client := redis.NewClient(config.GetPrimaryRedisOptions())

	// how to create a custom command
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

		key := "Project:45535"
		v, err := Get(client, key).Result()
		fmt.Printf("result: %s, err: %s\n", v, err)

		/// -- or

		x, err := client.Get("bad-key").Result()
		if err == redis.Nil {
			fmt.Println("key does not exist, redis.Nil...")
		} else {
			fmt.Printf("result: %s, err: %s\n", x, err)
			fmt.Println(x)
		}

		/// -- keys

		keys, err := client.Keys("Customer:*").Result()
		if err != nil {
			fmt.Println(err)
		} else if err == redis.Nil {
			fmt.Println("no keys found...")
		} else {
			fmt.Printf("customer keys, length: %d\n", len(keys))
			fmt.Println(keys)
		}
	}
}
