package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/elvisferns/redis-caching/cache"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Insuffiecient parameeters. Provide either get, set, delete")
		os.Exit(1)
		// testing2
	}

	ctxRoot := context.Background()
	redisCache, err := cache.NewClient(ctxRoot)

	if err != nil {
		fmt.Println("Redis server could not be connected", err)
		os.Exit(2)
	}

	defer redisCache.Close()

	redisCmd := os.Args[1]
	redisCmd = strings.ToLower(redisCmd)

	switch redisCmd {
	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		key := getCmd.String("key", "", "key used")
		getCmd.Parse(os.Args[2:])
		val, err := redisCache.Get(ctxRoot, *key)
		if err != nil {
			fmt.Println(err.(cache.MyError).What)
			return
		}

		fmt.Println("Key = ", *key, ":: Value = ", val.Val)

	case "set":
		setCmd := flag.NewFlagSet("set", flag.ExitOnError)
		key := setCmd.String("key", "", "key")
		value := setCmd.String("value", "", "value")
		expiry := setCmd.Int("ex", 0, "ex")
		setCmd.Parse(os.Args[2:])

		if *key == "" {
			fmt.Println("key cannot be empty")
			return
		}

		val, err := redisCache.Set(ctxRoot, *key, *value, time.Duration(*expiry)*time.Second)
		if err != nil {
			fmt.Println("Redis error while setting data:: ", err.(cache.MyError).What)
			return
		}

		fmt.Println("Status Code:: ", val.Val)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		key := deleteCmd.String("key", "", "key")
		deleteCmd.Parse(os.Args[2:])
		val, err := redisCache.Delete(ctxRoot, *key)
		if err != nil {
			fmt.Println(err.(cache.MyError).What)
			return
		}

		fmt.Println("Status Code", val.Val)

	default:
		fmt.Println("Invialid Command")
	}
}
