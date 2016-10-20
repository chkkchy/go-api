package hoge

import (
	"fmt"
	"time"

	// log "github.com/golang/glog"
	"gopkg.in/redis.v5"
)

const KEY = "hogetime"

var client *redis.Client

func init() {
	// Client
	client = redis.NewClient(&redis.Options{
		Addr:     "ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:6379",
		Password: "", // no password Set
		DB:       0,  //use default DB
	})

	// Connection close
	// defer client.Close()
}

func SetHogetime() {
	val := time.Now().String()
	err := client.Set(KEY, val, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetHogetime() interface{} {
	res, err := client.Get(KEY).Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return res
}
