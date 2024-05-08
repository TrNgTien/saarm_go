package services

import (
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func CreateUser() {

}
func GetUsers(c echo.Context) error {
	// limit, offset, page := c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("page")
	//
	// limitInt, err := utilities.GetIntValue(limit)
	//
	// if err != nil {
	// 	return err
	// }
	//
	// offsetInt, err := utilities.GetIntValue(offset)
	//
	// if err != nil {
	// 	return err
	// }
	//
	// pageInt, err := utilities.GetIntValue(page)
	//
	// if err != nil {
	// 	return err
	// }
	// users := repositories.UserRepo(pg.DB).FindAllUsers(common.PaginationQuery{
	// 	Limit:  limitInt,
	// 	Offset: offsetInt,
	// 	Page:   pageInt,
	// })

	return utilities.R200(c, "users")
}

func GetUserByID(id int) error {
	return nil
	// return repositories.UserRepo(pg.DB).FindUserByID(id)
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
