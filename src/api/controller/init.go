package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func Init() {
	e := echo.New()

	// routing
	users := e.Group("/users")
	users.GET("", getUsers)
	users.GET("/:id", getUser)
	users.POST("", postUser)
	users.PUT("/:id", putUser)
	users.DELETE("/:id", deleteUser)

	e.Static("/static", "static")

	e.Run(standard.New(":1323"))
}
