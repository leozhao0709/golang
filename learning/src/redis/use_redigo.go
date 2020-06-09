package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func redigoSetGetHash() {
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis connection err", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("hmset", "user1", "name", "john", "age", 30)
	if err != nil {
		fmt.Println("hmset error", err)
		return
	}

	val1, err := redis.StringMap(conn.Do("hgetall", "user1"))
	if err != nil {
		fmt.Println("hgetall fail", err)
		return
	}
	fmt.Println(val1)

	val2, err := redis.Strings(conn.Do("hmget", "user1", "name", "age"))
	if err != nil {
		fmt.Println("hgetall fail", err)
		return
	}
	fmt.Println(val2)

	_, err = conn.Do("expire", "user1", 30)
	if err != nil {
		fmt.Println("set expire err", err)
		return
	}
}

func redigoSetGetString() {
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis connection err", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("set", "name", "tomjerry1")
	if err != nil {
		fmt.Println("set err", err)
		return
	}

	name, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err", err)
		return
	}
	fmt.Println("get name:", name)

	_, err = conn.Do("expire", "name", 30)
	if err != nil {
		fmt.Println("set expire err", err)
		return
	}
}
