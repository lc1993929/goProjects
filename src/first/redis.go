package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}
	defer func(c redis.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(c)

	//err = c.Send("set", "123", "123")
	_, err = c.Do("set", "123", "123")

	if err != nil {
		fmt.Println(err)
	}
	//err = c.Send("get", "123")
	//result, err := c.Do("get", "123")
	//fmt.Println(string(result.([]byte)))
	result, err := redis.String(c.Do("get", "123"))
	fmt.Println(result)
	/*	if err != nil {
			fmt.Println(err)
		}
		err = c.Flush()
		if err != nil {
			fmt.Println(err)
		}
		result, err := c.Receive()
		if err != nil {
			fmt.Println(err)
		}*/

}
