package controllers

import (
	"github.com/labstack/echo/v4"
)

func GetApartments(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetApartmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func PutApartmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func PatchApartmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteApartmentByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
