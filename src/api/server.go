package main

import (
	"api/controller"
	"core/dao"
	"core/logger"
)

func main() {
	logger.Init()
	dao.Init()
	controller.Init()
}
