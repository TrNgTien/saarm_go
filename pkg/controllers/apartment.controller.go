package controllers

import (
	modelRequests "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func CreateApartments(c echo.Context) error {
	a := new(modelRequests.NewApartment)

	if err := c.Bind(a); err != nil {
		return utilities.R400(c, err.Error())
	}

	userID := c.Get("userID").(string)

	apartment := modelRequests.NewApartment{
		Name:          a.Name,
		LocationUrl:   a.LocationUrl,
		Address:       a.Address,
		TotalRoom:     a.TotalRoom,
		RoomAvailable: a.RoomAvailable,
		UserID:        utilities.ParseStringToUuid(userID),
	}

	apartmentCreated, err := services.CreateApartments(apartment)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, apartmentCreated)
}

func GetApartments(c echo.Context) error {
	apartments, err := services.GetApartments()

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, apartments)
}

func GetApartmentsByUserID(c echo.Context) error {
	ID := c.Param("id")
	userID := utilities.ParseStringToUuid(ID)

	apartments, err := services.GetApartmentsByUserID(userID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, apartments)
}

func GetApartmentByID(c echo.Context) error {
	ID := c.Param("id")
	apartmentId := utilities.ParseStringToUuid(ID)
	apartment, err := services.GetApartmentByID(apartmentId)

	if err != nil {
		return utilities.R400(c, "[GetAparments] Cannot get apartment!")
	}

	return utilities.R200(c, apartment)
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
