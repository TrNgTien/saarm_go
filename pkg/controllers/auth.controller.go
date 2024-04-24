package controllers

import (
	"fmt"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"saarm/pkg/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

// const (
// 	clientID     = "YOUR_CLIENT_ID"     // Replace with your actual client ID
// 	clientSecret = "YOUR_CLIENT_SECRET" // Replace with your actual client secret
// 	redirectURI  = "postmessage"
// 	grantType    = "authorization_code"
// )

// type TokenResponse struct {
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// 	// Add other relevant fields from the response here
// }


func SignIn(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}

	user := common.UserDTO{
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
	u := new(User)
	if err := c.Bind(u); err != nil {
		return utilities.R400(c, "err")
	}

	user := common.UserDTO{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return utilities.R400(c, "[SignUp] | Missing request fields")
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
