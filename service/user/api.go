package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (_u *User) GetUsers(c echo.Context) error {
	sex := c.QueryParam("sex")

	ret := _u.findUsers(sex)

	return c.JSON(http.StatusOK, ret)
}

func (_u *User) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := _u.findUser(id)

	return c.JSON(http.StatusOK, ret)
}

func (_u *User) PostUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := _u.createUser(u)

	return c.JSON(http.StatusCreated, ret)
}

func (_u *User) PutUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := _u.updateUser(id, u.Name, u.Age)

	return c.JSON(http.StatusOK, ret)
}

func (_u *User) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ret := _u.deleteUser(id)

	return c.JSON(http.StatusNoContent, ret)
}
