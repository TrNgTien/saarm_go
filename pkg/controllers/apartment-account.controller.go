package controllers

import (
	"fmt"
	"saarm/pkg/common"
	modelRequests "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func CreateAccount(c echo.Context) error {
	a := new(modelRequests.NewApartment)

	if err := c.Bind(a); err != nil {
		return utilities.R400(c, err.Error())
	}

	apartment := modelRequests.NewApartment{
		Name:          a.Name,
		LocationUrl:   a.LocationUrl,
		Address:       a.Address,
		TotalRoom:     a.TotalRoom,
		RoomAvailable: a.RoomAvailable,
	}

	apartmentCreated, err := services.CreateApartments(apartment)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, apartmentCreated)
}


func GetAccount(c echo.Context) error {

	limit, offset, page := c.QueryParam("limit"), c.QueryParam("offset"), c.QueryParam("page")
	if limit == "" || offset == "" || page == "" {
		return utilities.R400(c, "[GetAparments] TODO Paging!")
	}

	limitInt, err := utilities.GetIntValue(limit)

	if err != nil {
		return err
	}

	offsetInt, err := utilities.GetIntValue(offset)

	if err != nil {
		return err
	}

	pageInt, err := utilities.GetIntValue(page)

	if err != nil {
		return err
	}

	queryData := common.PaginationQuery{
		Limit:  limitInt,
		Offset: offsetInt,
		Page:   pageInt,
	}

	fmt.Println("[GetAparments] runign here")
	apartments, err := services.GetApartments(queryData)

	if err != nil {
		return utilities.R400(c, "[GetAparments] Cannot get apartments!")
	}

	return utilities.R200(c, apartments)
}

func GetAccountByID(c echo.Context) error {
	ID := c.Param("id")
	apartmentId := utilities.ParseStringToUuid(ID)
	apartment, err := services.GetAparmentByID(apartmentId)

	if err != nil {
		return utilities.R400(c, "[GetAparments] Cannot get apartment!")
	}

	return utilities.R200(c, apartment)
}

func UpdateAccount(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}


func DeleteAcount(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteAcountByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}
