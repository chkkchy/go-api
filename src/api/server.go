package main

import (
	"flag"

	"api/controller"

	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {

	// logger
	flag.Parse()
	defer glog.Flush()
	glog.Info("hig")

	e := echo.New()

	// routing
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.PostUser)
	e.PUT("/users/:id", controller.PutUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.Static("/static", "static")

	e.Run(standard.New(":1323"))

}
