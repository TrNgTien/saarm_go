package services

import (
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/models"
	"saarm/pkg/repositories"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func GetAparments(c echo.Context) error {
	limit, offset, page := c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("page")
	limitInt, err := utilities.GetIntValue(limit)

	if err != nil {
		return err
	}

	offsetInt, err := utilities.GetIntValue(offset)

	if err != nil {
		return err
	}

	pageInt, err := utilities.GetIntValue(page)

	if err != nil {
		return err
	}

	users := repositories.AparmentRepo(pg.DB).FindAllAparments(common.PaginationQuery{
		Limit:  limitInt,
		Offset: offsetInt,
		Page:   pageInt,
	})

	return utilities.R200(c, users)
}

func GetAparmentByID(id int) models.AparmentResponse {
	return repositories.AparmentRepo(pg.DB).FindAparmentByID(id)
}

func PatchAparment(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func PutAparments(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteAparmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}

func DeleteAparments(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    "users",
	})
}
