package controllers

import (
	modelRequest "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func CreateRole(c echo.Context) error {

	r := new(modelRequest.NewRole)

	if err := c.Bind(r); err != nil {
		return echo.ErrBadRequest
	}

  if r.Name == "" {

		return utilities.R400(c, "[CreateRole] Please provide role name!")
  }
	role := modelRequest.NewRole{
		Name: r.Name,
	}

	createRoleErr := services.CreateRole(role)

	if createRoleErr != nil {
		return utilities.R400(c, "[CreateRole] Cannot create role!"+createRoleErr.Error())
	}

	return utilities.R204(c)
}

func GetRoles(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func UpdateRoles(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteRoles(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
