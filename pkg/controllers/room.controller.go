package controllers

import (
	"encoding/base64"
	"os"
	"saarm/pkg/utilities"
	"strings"

	"github.com/labstack/echo/v4"
)

type UploadWaterMeter struct {
	File string `json:"file"`
}

func GetRooms(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
}

func GetRoomByID(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"success": true,
	})
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

func GetWaterMeter(c echo.Context) error {

	file := new(UploadWaterMeter)

	if err := c.Bind(file); err != nil {
		return utilities.R400(c, err.Error())
	}

  baseData := file.File[strings.IndexByte(file.File, ',')+1:]

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		return utilities.R500(c, err.Error())
	}

	f, err := os.Create("Uploadedfil.png")
	if err != nil {
		return utilities.R500(c, err.Error())
	}
	defer f.Close()

	if _, err := f.Write(decodedBase64); err != nil {
		return utilities.R500(c, err.Error())
	}
	if err := f.Sync(); err != nil {
		return utilities.R500(c, err.Error())
	}

	// texts, err := services.SubmitWaterMeter(file)

	// if err != nil {
	// 	return utilities.R400(c, err.Error())
	// }

	return utilities.R200(c, "data")
}
