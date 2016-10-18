package main

import (
	"flag"

	userService "./service/user"

	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

var user userService.User

func main() {

	// logger
	flag.Parse()
	defer glog.Flush()

	e := echo.New()

	// routing
	e.GET("/users", user.GetUsers)
	e.GET("/users/:id", user.GetUser)
	e.POST("/users", user.PostUser)
	e.PUT("/users/:id", user.PutUser)
	e.DELETE("/users/:id", user.DeleteUser)

	e.Static("/static", "static")

	e.Run(standard.New(":1323"))

}
