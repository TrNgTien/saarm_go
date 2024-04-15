package services

import "github.com/labstack/echo/v4"

func GetUsers(c echo.Context) error {
  	return c.JSON(200, echo.Map{
		"success": true,
    "data": "users",
	})
}

func GetUserByID(id int) int {
  return id;
}

func PatchUser(c echo.Context) error {
  	return c.JSON(200, echo.Map{
		"success": true,
    "data": "users",
	})
}

func PutUsers(c echo.Context) error {
  	return c.JSON(200, echo.Map{
		"success": true,
    "data": "users",
	})
}

func DeleteUserByID(c echo.Context) error {
  	return c.JSON(200, echo.Map{
		"success": true,
    "data": "users",
	})
}

func DeleteUsers(c echo.Context) error {
  	return c.JSON(200, echo.Map{
		"success": true,
    "data": "users",
	})
}
