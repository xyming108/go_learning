package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"math/rand"
	"strconv"
)

func main() {
	dial, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("redis连接失败", err)
		return
	}
	defer dial.Close()

	for i := 0; i < 100; i++ {
		dial.Do("zadd", "rank", rand.Intn(1000), "Kevin-"+strconv.Itoa(i))
	}

	slices, err := redis.ByteSlices(dial.Do("zrevrange", "rank", 0, -1, "withscores"))
	for _, v := range slices {
		fmt.Println(string(v))
	}
}
