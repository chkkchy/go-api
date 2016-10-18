package user

import (
	"net/http"
	"strconv"

	m "../model"
	userService "../service"
	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex")

	ret := userService.FindUsers(sex)

	return c.JSON(http.StatusOK, ret)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := userService.FindUser(id)

	return c.JSON(http.StatusOK, ret)
}

func PostUser(c echo.Context) error {
	u := new(m.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := userService.CreateUser(u)

	return c.JSON(http.StatusCreated, ret)
}

func PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	u := new(m.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := userService.UpdateUser(id, u.Name, u.Age)

	return c.JSON(http.StatusOK, ret)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := userService.DeleteUser(id)

	return c.JSON(http.StatusNoContent, ret)
}
