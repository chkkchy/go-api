package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func Init() {
	e := echo.New()

	// routing
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", postUser)
	e.PUT("/users/:id", putUser)
	e.DELETE("/users/:id", deleteUser)

	e.Static("/static", "static")

	e.Run(standard.New(":1323"))
}
