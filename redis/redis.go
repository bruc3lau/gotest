package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println(rdb)

	rdb.Set("a", "1", time.Second*10).Err()

	a := rdb.Get("a")
	fmt.Println(a)

	rdb.SetNX("b", 2, time.Hour)
	fmt.Println(rdb.Get("b"))
}
