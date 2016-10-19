package controller

import (
	"net/http"
	"strconv"

	"../service/user"
	"github.com/labstack/echo"
)

var service *user.Service = user.New()

func GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex") == "male"
	req := &user.GetUsersInput{Sex: sex}
	res := service.GetUsers(req)
	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := service.GetUser(id)
	return c.JSON(http.StatusOK, res)
}

func PostUser(c echo.Context) error {
	req := new(user.CreateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := service.CreateUser(req)
	return c.JSON(http.StatusCreated, res)
}

func PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(user.UpdateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req.Id = id
	res := service.UpdateUser(req)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := service.DeleteUser(id)
	return c.JSON(http.StatusNoContent, res)
}
