package user

import (
	"net/http"
	"strconv"

	"./useriface"
	"github.com/labstack/echo"
)

var Api useriface.UserAPI = &Service{}

func (service *Service) GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex")
	ret := service.findUsers(sex)
	return c.JSON(http.StatusOK, ret)
}

func (service *Service) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ret := service.findUser(id)
	return c.JSON(http.StatusOK, ret)
}

func (service *Service) PostUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ret := service.createUser(user)
	return c.JSON(http.StatusCreated, ret)
}

func (service *Service) PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ret := service.updateUser(id, user.Name, user.Age)
	return c.JSON(http.StatusOK, ret)
}

func (service *Service) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ret := service.deleteUser(id)
	return c.JSON(http.StatusNoContent, ret)
}
