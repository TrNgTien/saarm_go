package controllers

import (
	"fmt"
	modelRequest "saarm/pkg/models/request"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"saarm/pkg/services"
	"saarm/pkg/utilities"
)

func GetUsers(c echo.Context) error {
	return services.GetUsers(c)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	userID := utilities.ParseStringToUuid(id)

	user, err := services.GetUserByID(userID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, user)
}

func CreateUser(c echo.Context) error {
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

	userData, err := services.CreateUser(user)

	if err != nil {
		fmt.Println("SignUp ", err)
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, userData)
}

func UpdateUser(c echo.Context) error {
	return nil
}

func PatchUser(c echo.Context) error {
	id := c.Param("id")
	userId := utilities.ParseStringToUuid(id)

	req := new(modelRequest.UpdateUserRequest)

	if err := c.Bind(req); err != nil {
		return utilities.R400(c, "Bad Request!")
	}

	userReq := modelRequest.UpdateUserRequest{
		Username: req.Username,
		Password: req.Password,
		RoleName: req.RoleName,
		Email:    req.Email,
		Status:   req.Status,
	}

	if err := services.PatchUser(userId, userReq); err != nil {
		return utilities.R400(c, "Cannot update user!!")
	}

	return utilities.R204(c)
}
