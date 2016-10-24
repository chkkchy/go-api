package cache

import (
	"time"

	log "github.com/golang/glog"
	redis "gopkg.in/redis.v5"
)

const KEY = "hogetime"

func SetHogetime() {
	val := time.Now().String()
	err := client.Set(KEY, val, 0).Err()
	if err != nil {
		log.Error(err.Error())
	}
}

func GetHogetime() interface{} {
	res, err := client.Get(KEY).Result()
	if err == redis.Nil {
		log.Info(KEY + "does not exists")
	} else if err != nil {
		log.Error(err.Error())
	}
	return res
}
