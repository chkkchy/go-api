package test

import (
	"fmt"
	"os"

	// log "github.com/golang/glog"
	"gopkg.in/redis.v5"
)

func init() {

	fmt.Println("goredis init")

	client := redis.NewClient(&redis.Options{
		Addr:     "ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:6379",
		Password: "", // no password Set
		DB:       0,  //use default DB
	})

	// Connection close
	defer client.Close()

	const key = "keygoredis"
	const val = "valgoredis"

	// Set
	client.Set(key, val, 0).Err()

	// Get
	res, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v", res)

}
