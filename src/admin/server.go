package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {

	e := echo.New()

	// routing
	e.GET("/", test)

	e.Run(standard.New(":1324"))

}

func test(c echo.Context) error {
	fmt.Println("[admin] test!!")
	return c.String(http.StatusOK, "admin response!!")
}
