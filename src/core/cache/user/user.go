package user

import (
	"fmt"
	"os"

	// log "github.com/golang/glog"
	"github.com/garyburd/redigo/redis"
)

func init() {

	fmt.Println("redis init")

	c, err := redis.Dial("tcp", "ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:6379")
	if err != nil {
		fmt.Println("dial err")
		fmt.Println(err)
		os.Exit(1)
	}

	// Connection close
	defer c.Close()

	const key = "hoge"
	const val = "fugo"

	// Set
	c.Do("SET", key, val)

	// Get
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v", s)

}
