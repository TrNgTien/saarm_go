package controllers

import (
	"fmt"
	modelRequest "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SignIn(c echo.Context) error {
	u := new(modelRequest.SignInRequest)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user := modelRequest.SignInRequest{
		Username: u.Username,
		Password: u.Password,
	}

	userData, err := services.SignIn(user)

	if err != nil {
		return utilities.R400(c, "[SignIn] Cannot get user!")
	}

	return utilities.R200(c, userData)
}

func SignUp(c echo.Context) error {
	u := new(modelRequest.NewUser)
	if err := c.Bind(u); err != nil {
		return utilities.R400(c, err.Error())
	}

	user := modelRequest.SignUpRequest{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	if user.Username == "" || user.Password == "" {
		return utilities.R400(c, "[SignUp] | Missing username or password!")
	}

	if services.IsExistedUser(user) {
		log.Error("[SignUp] | User has already exsited!")
		return utilities.R400(c, "[SignUp] | User has already exsited!")
	}

	userData, err := services.SignUp(user)

	if err != nil {
		fmt.Println("SignUp ", err)
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, userData)
}
