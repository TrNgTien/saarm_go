package services

import (
	"saarm/modules/pg"
	"saarm/pkg/common"
	"saarm/pkg/utilities"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Result struct {
	id uuid.UUID
}

func isExistedUser(user common.UserDTO) bool {
	var count int
	row := pg.DB.Raw("select count(*) from users where username = ?", user.Username).Row()

	row.Scan(&count)

	if count > 0 {
		return true
	}

	return false
}

func SignIn(user common.UserDTO) error {
	isExistedUser(user)
	return echo.ErrBadRequest
}

func SignUp(c echo.Context, user common.UserDTO) error {

	if isExistedUser(user) {
		log.Error("[SignUp] | User has already exsited!")

    return utilities.R400(c, "[SignUp] | User has already exsited!")
	}


  return utilities.R200(c, "data")
}
