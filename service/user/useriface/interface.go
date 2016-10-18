package useriface

import "github.com/labstack/echo"

type UserAPI interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	PostUser(c echo.Context) error
	PutUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

// var _ UserAPI = (*user.User)(nil)
