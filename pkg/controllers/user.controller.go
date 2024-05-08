package controllers

import (
	"github.com/labstack/echo/v4"

	"saarm/pkg/services"
	"saarm/pkg/utilities"
)

func GetUsers(c echo.Context) error {
	return services.GetUsers(c)
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	i, err := utilities.GetIntValue(id)
	if err != nil {

		return utilities.R400(c, err.Error())

	}

	rs := services.GetUserByID(i)

	return utilities.R200(c, rs)
}


func CreateUser(c echo.Context) error {
	return utilities.R200(c, "rs")
}

func UpdateUser(c echo.Context) error {
	return utilities.R200(c, "rs")
}
