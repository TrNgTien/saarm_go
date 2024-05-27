package utilities

import "github.com/labstack/echo/v4"

func R204(c echo.Context) error {
	return c.JSON(204, echo.Map{
		"success": true,
	})
}

func R200[T any](c echo.Context, data T) error {
	return c.JSON(200, echo.Map{
		"success": true,
		"data":    data,
	})
}

func R400(c echo.Context, msg string) error {
	return c.JSON(400, echo.Map{
		"success": false,
		"message": msg,
	})
}

func R401(c echo.Context, msg string) error {
	return c.JSON(401, echo.Map{
		"success": false,
		"message": "Unauthorized!",
	})
}

func R403(c echo.Context) error {
	return c.JSON(403, echo.Map{
		"success": false,
		"message": "Forbidden",
	})
}

func R500(c echo.Context, msg string) error {
	return c.JSON(500, echo.Map{
		"success": false,
		"message": msg,
	})
}
