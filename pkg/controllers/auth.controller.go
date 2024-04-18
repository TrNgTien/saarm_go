package controllers

import (
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"saarm/pkg/common"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func SignIn(c echo.Context) (err error) {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user := common.UserDTO{
		Username: u.Username,
		Password: u.Password,
	}

	userData := services.SignIn(user)

	if userData == nil {
		return utilities.R400(c, "[SignIn] Cannot get user!")
	}

	return utilities.R204(c)
}

func SignUp(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user := common.UserDTO{
		Username: u.Username,
		Password: u.Password,
	}

	return services.SignUp(c, user)
}
