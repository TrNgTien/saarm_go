package controllers

import (
	"fmt"
	modelRequests "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SignIn(c echo.Context) error {
	u := new(modelRequests.SignInRequest)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user := modelRequests.SignInRequest{
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
	u := new(modelRequests.NewUser)
	if err := c.Bind(u); err != nil {
		return utilities.R400(c, err.Error())
	}

	user := modelRequests.SignUpRequest{
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		ApartmentID: u.ApartmentID,
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return utilities.R400(c, "[SignUp] | Missing request fields")
	}

	if services.IsExistedUser(user) {
		fmt.Println("runn")
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
