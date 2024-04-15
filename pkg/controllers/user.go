package controllers

import (
  "strconv"

  "github.com/labstack/echo/v4"

  "saarm/pkg/services"
  "saarm/pkg/utilities"
)

func GetUsers (c echo.Context) error {
  return services.GetUsers(c)
}

func GetUserByID (c echo.Context) error {
  id := c.Param("id")
  i, err := strconv.Atoi(id)

  if err != nil {
    return utilities.R400(c, "[GetUserByID] Cannot parse to int" )
  }

  rs := services.GetUserByID(i)


  return c.JSON(200, echo.Map{
    "success": true,
    "data": rs,
  })
}
