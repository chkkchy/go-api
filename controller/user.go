package user

import (
	"net/http"
	"strconv"

	userService "../service"
	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"bool"` // false -> male, true -> female
}

func GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex")

	ret := userService.FindUsers(sex)

	return c.JSON(http.StatusOK, ret)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ret := userService.FindUser(id)

	return c.JSON(http.StatusOK, ret)
}

func PostUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	ret := userService.CreateUser(u)

	return c.JSON(http.StatusCreated, ret)
}

func PutUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	ret := userService.UpdateUser(id, u.Name, u.Age)

	return c.JSON(http.StatusOK, ret)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ret := userService.DeleteUser(id)

	return c.JSON(http.StatusOK, ret)
}
