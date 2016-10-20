package controller

import (
	"net/http"
	"strconv"

	"core"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex") == "male"
	req := &core.GetUsersInput{Sex: sex}
	res := core.GetUsers(req)
	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := core.GetUser(id)
	return c.JSON(http.StatusOK, res)
}

func PostUser(c echo.Context) error {
	req := new(core.CreateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := core.CreateUser(req)
	return c.JSON(http.StatusCreated, res)
}

func PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(core.UpdateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req.Id = id
	res := core.UpdateUser(req)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := core.DeleteUser(id)
	return c.JSON(http.StatusNoContent, res)
}
