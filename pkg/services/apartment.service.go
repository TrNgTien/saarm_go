package services

import (
  "saarm/modules/pg"
  "saarm/pkg/models"
  "saarm/pkg/repositories"
  "saarm/pkg/utilities"

  "github.com/labstack/echo/v4"
)

func GetAparments(c echo.Context) error {
  limit, offset, page := c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("page")

  users := repositories.AparmentRepo(pg.DB).FindAllAparments(repositories.PaginationQuery{
    Limit: utilities.GetIntValue(limit),
    Offset: utilities.GetIntValue(offset),
    Page: utilities.GetIntValue(page),
  })

  return utilities.R200(c, users)
}

func GetAparmentByID(id int) models.AparmentResponse {
  return repositories.AparmentRepo(pg.DB).FindAparmentByID(id)
}

func PatchAparment(c echo.Context) error {
  return c.JSON(200, echo.Map{
    "success": true,
    "data": "users",
  })
}

func PutAparments(c echo.Context) error {
  return c.JSON(200, echo.Map{
    "success": true,
    "data": "users",
  })
}

func DeleteAparmentByID(c echo.Context) error {
  return c.JSON(200, echo.Map{
    "success": true,
    "data": "users",
  })
}

func DeleteAparments(c echo.Context) error {
  return c.JSON(200, echo.Map{
    "success": true,
    "data": "users",
  })
}
