package main

import (
	"api/controller"
	"core/cache"
	"core/dao"
	"core/logger"
)

func main() {
	logger.Init()
	dao.Init()
	cache.Init()
	controller.Init()
}
