package utilities

import "github.com/labstack/echo/v4"

func R200(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func R400(c echo.Context, msg string) error {
	return c.JSON(400, echo.Map{
		"success": false,
		"message": msg,
	})
}
