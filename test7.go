package main

import (
	"github.com/go-redis/redis"
)

func main() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if nx := redis.SetNX("lock_key", "lock_value", 0); nx.Val() == true {
		redis.Expire("lock_key", 100)

		// todo 执行业务逻辑

		redis.Del("lock_key")
	}

	if nx := redis.SetNX("lock_key", "lock_value", 1000); nx.Val() == true {
		// todo 执行业务逻辑

		if redis.Get("lock_key").Val() == "lock_value" {
			redis.Del("lock_key")
		}
	}
}
