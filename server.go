package main

import (
	"flag"

	"./service/user"

	"github.com/golang/glog"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {

	// logger
	flag.Parse()
	defer glog.Flush()

	e := echo.New()

	// routing
	e.GET("/users", user.Api.GetUsers)
	e.GET("/users/:id", user.Api.GetUser)
	e.POST("/users", user.Api.PostUser)
	e.PUT("/users/:id", user.Api.PutUser)
	e.DELETE("/users/:id", user.Api.DeleteUser)

	e.Static("/static", "static")

	e.Run(standard.New(":1323"))

}
