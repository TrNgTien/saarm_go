package controllers

import (
	modelRequest "saarm/pkg/models/request"

	"github.com/labstack/echo/v4"

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
	return utilities.R200(c, "rs")
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
