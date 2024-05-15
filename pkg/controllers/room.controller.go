package controllers

import (
	"saarm/pkg/common"
	modelRequest "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
)

func CreateRoom(c echo.Context) error {
	r := new(modelRequest.NewRoom)

	if err := c.Bind(r); err != nil {
		return utilities.R400(c, err.Error())
	}

	room := modelRequest.NewRoom{
		Username:      r.Username,
		Password:      r.Password,
		Status:        r.Status,
		Name:          r.Name,
		RoomPrice:     r.RoomPrice,
		MaxPeople:     r.MaxPeople,
		CurrentPeople: r.CurrentPeople,
		ApartmentID:   r.ApartmentID,
	}

	if room.Username == "" || room.Password == "" {
		return utilities.R400(c, "[CreateRoom] | Missing username or password!")
	}

	roomCreated, err := services.CreateRoom(room)

	if err != nil {
		return utilities.R400(c, "[CreateRoom] | "+err.Error())
	}

	return utilities.R200(c, roomCreated)
}

func GetRooms(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetRoomByID(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)
	return services.GetRoomByID(roomID)
}

func PutRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func PatchRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func LoginRoom(c echo.Context) error {

	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetWaterMeter(c echo.Context) error {
	roomID := c.Param("id")
	file := new(common.UploadWaterMeter)

	if err := c.Bind(file); err != nil {
		return utilities.R400(c, err.Error())
	}

	numberDetected, err := services.SubmitWaterMeter(*file, roomID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, numberDetected)
}

func DuplicateRoom(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)
	duplicatedRoom, err := services.DuplicateRoom(roomID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, duplicatedRoom)
}
