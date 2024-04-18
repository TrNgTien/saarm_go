package services

import (
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/models"
	"saarm/pkg/repositories"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	limit, offset, page := c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("page")

	users := repositories.UserRepo(pg.DB).FindAllUsers(common.PaginationQuery{
		Limit:  utilities.GetIntValue(limit),
		Offset: utilities.GetIntValue(offset),
		Page:   utilities.GetIntValue(page),
	})

	return utilities.R200(c, users)
}

func GetUserByID(id int) models.UserResponse {
	return repositories.UserRepo(pg.DB).FindUserByID(id)
}

func PatchUser(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func PutUsers(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteUserByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteUsers(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}
