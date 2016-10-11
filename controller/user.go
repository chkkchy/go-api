package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"bool"` // false -> male, true -> female
}

var users map[int]User = map[int]User{
	1: User{"yukichi", 26, false},
	2: User{"mendy", 25, true},
	3: User{"shibata", 28, false},
}

func GetUsers(c echo.Context) error {
	_ := c.QueryParam("sex")

	ret := mToS(users)
	return c.JSON(http.StatusOK, ret)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, users[id])
}

func PostUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	users[(len(users) + 1)] = *u

	return c.JSON(http.StatusCreated, u)
}

func PutUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user := users[id]
	user.Name = u.Name
	user.Age = u.Age
	users[id] = user

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(users, id)

	_, ok := users[id]
	deleted := !ok

	return c.JSON(http.StatusOK, deleted)
}

func mToS(m map[int]User) []User {
	v := make([]User, 0, len(m))
	for _, value := range m {
		v = append(v, value)
	}
	return v
}
