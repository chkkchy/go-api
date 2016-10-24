package cache

import (
	"github.com/golang/glog"
	redis "gopkg.in/redis.v5"
)

var client *redis.Client

func Init() {
	// Client
	client = redis.NewClient(&redis.Options{
		Addr:     "ec2-52-198-155-206.ap-northeast-1.compute.amazonaws.com:6379",
		Password: "", // no password Set
		DB:       0,  //use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		glog.Fatal(err.Error())
	}

	// Connection close
	// defer client.Close()
}
