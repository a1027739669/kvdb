package main

import "github.com/garyburd/redigo/redis"

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic("error connecting")
	}
	defer conn.Close()
	conn.Do("SET", "abc", 123)
}
