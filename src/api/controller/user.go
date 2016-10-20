package controller

import (
	"net/http"
	"strconv"

	"core/v1"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex") == "male"
	req := &v1.GetUsersInput{Sex: sex}
	res := v1.GetUsers(req)
	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	// var ha string = c.Request().Header().Get("Accept")
	// glog.Info("header[accept]", ha)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := v1.GetUser(id)
	return c.JSON(http.StatusOK, res)
}

func PostUser(c echo.Context) error {
	req := new(v1.CreateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := v1.CreateUser(req)
	return c.JSON(http.StatusCreated, res)
}

func PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req := new(v1.UpdateUserInput)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	req.Id = id
	res := v1.UpdateUser(req)
	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res := v1.DeleteUser(id)
	return c.JSON(http.StatusNoContent, res)
}
