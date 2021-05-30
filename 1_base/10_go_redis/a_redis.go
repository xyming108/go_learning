package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/**
Author: xym
Date: 2021/5/9 21:35
Project: myproject2_go_learning
Description:
*/

/*
Redis运行在内存中但是可以持久化到磁盘
*/

//String
func s(r redis.Conn) {
	_, err := r.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr, err := redis.Int(r.Do("Get", "abc"))
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println("获取成功", rr)

	//********String的批量操作*****************
	_, err = r.Do("MSet", "a", 1, "b", 2)
	if err != nil {
		fmt.Println("批量操作String失败", err)
		return
	}
	mrr, err := redis.Ints(r.Do("MGet", "a", "b"))
	if err != nil {
		fmt.Println("批量操作String失败", err)
		return
	}

	fmt.Println("批量获取成功", mrr)
}

//List
func l(r redis.Conn) {
	_, err := r.Do("flushdb")
	if err != nil {
		fmt.Println("清除错误", err)
		return
	}
	_, err = r.Do("lpush", "list", "hh", "ee", "mm")
	if err != nil {
		fmt.Println("队列插入失败", err)
		return
	}

	//lpop, err := redis.String(r.Do("lrange list", "list"))
	values, err := redis.Strings(r.Do("lrange", "list", 0, -1))
	if err != nil {
		fmt.Println("队列获取失败", err)
		return
	}

	fmt.Println("队列结果获取成功", values)
}

//hash
func h(r redis.Conn) {
	_, err := r.Do("HSet", "hbook", "hhh", "123")
	if err != nil {
		fmt.Println("hash失败", err)
		return
	}
	h, err := redis.Int(r.Do("HGet", "hbook", "hhh"))
	if err != nil {
		fmt.Println("失败", err)
		return
	}
	fmt.Println("从hash中获取成功", h)

}

func main() {
	//连接之前要打开服务器（redis-server）
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis连接失败", err)
		return
	}
	fmt.Println("redis连接成功")
	defer r.Close()
	s(r)
	fmt.Println()
	l(r)
	fmt.Println()
	h(r)
}
