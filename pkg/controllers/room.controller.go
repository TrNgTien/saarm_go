package controllers

import (
	"github.com/labstack/echo/v4"
)

func GetRooms(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func PutRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func PatchRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
