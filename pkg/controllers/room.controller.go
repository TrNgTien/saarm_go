package controllers

import (
	"fmt"
	"saarm/pkg/common"
	modelRequest "saarm/pkg/models/request"
	"saarm/pkg/services"
	"saarm/pkg/utilities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CreateRoom(c echo.Context) error {
	r := new(modelRequest.NewRoom)

	if err := c.Bind(r); err != nil {
		return utilities.R400(c, err.Error())
	}

	fmt.Println("[CreateRoom]", r.CurrentPeople)
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

	if services.IsExistedRoomAccount(room) {
		log.Error("[CreateRoom] | Room Account has already exsited!")
		return utilities.R400(c, "Tài khoản đã tồn tại!")
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
	roomByID, err := services.GetRoomByID(roomID)

	if err != nil {
		return utilities.R400(c, "[GetRoomByID] | "+err.Error())

	}

	return utilities.R200(c, roomByID)
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

func DeleteRooms(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DeleteRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func DetectWaterMeter(c echo.Context) error {
	roomID := c.Param("id")
	file := new(common.UploadWaterMeter)

	if err := c.Bind(file); err != nil {
		fmt.Println("[GetWaterMeter][Bind body]: ", err.Error())
		return utilities.R400(c, err.Error())
	}

	numberDetected, err := services.DetectWaterMeter(*file, roomID)

	if err != nil {
		fmt.Println("[GetWaterMeter][detection service]: ", err.Error())
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, numberDetected)
}

func GetBillByRoom(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)
	monthRequest := c.QueryParam("monthRequest")

	data, err := services.GetBillByRoom(roomID, monthRequest)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, data)
}

func GetBills(c echo.Context) error {

	data, err := services.GetBills()

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, data)
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

func ConfirmWaterMeter(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)

	waterMeterNumber := new(modelRequest.SubmitWaterMeterNumber)

	if err := c.Bind(waterMeterNumber); err != nil {
		fmt.Println("[ConfirmWaterMeter][Bind body]: ", err.Error())
		return utilities.R400(c, err.Error())
	}

	err := services.ConfirmWaterMeter(roomID, waterMeterNumber.WaterMeter)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, "Submitted Successfully!!")
}

func CheckSubmittedWaterMeter(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)
	isSubmitted, err := services.CheckSubmittedWaterMeter(roomID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, isSubmitted)
}

func GetHistorySubmitted(c echo.Context) error {
	ID := c.Param("id")
	roomID := utilities.ParseStringToUuid(ID)
	histories, err := services.GetHistorySubmitted(roomID)

	if err != nil {
		return utilities.R400(c, err.Error())
	}

	return utilities.R200(c, histories)
}
