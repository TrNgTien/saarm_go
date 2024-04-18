package controllers

import (
	"github.com/labstack/echo/v4"
)

func GetWaterMeter(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
