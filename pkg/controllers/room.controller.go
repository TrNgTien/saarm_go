package controllers

import (
	"saarm/pkg/common"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

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

func LoginRoom(c echo.Context) error {

	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetWaterMeter(c echo.Context) error {
	roomID := c.Param("id")
	file := new(common.UploadWaterMeter)

	if err := c.Bind(file); err != nil {
		return utilities.R400(c, err.Error())
	}

	numberDetected, err := services.SubmitWaterMeter(*file, roomID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, numberDetected)
}
