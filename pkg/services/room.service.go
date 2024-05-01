package services

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"saarm/pkg/common"
	"saarm/pkg/helpers"
	"saarm/pkg/utilities"
	"strings"
)

func StoreWaterMeterFile(file common.UploadWaterMeter, roomID string) error {
	baseData := file.File[strings.IndexByte(file.File, ',')+1:]

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		return err
	}

	parts := strings.SplitN(file.File, ";", 2)
	var fileType string

	if len(parts) != 2 {
		return err
	}

	mimeType := parts[0]
	fileType = strings.Split(mimeType, "/")[1]

	timestamp := helpers.GetCurrentTimestampString()
	outputFileName := "room-" + roomID + "-" + timestamp + "." + fileType

	f, err := os.Create(utilities.GetFilePath(common.WATER_METER_PATH + outputFileName))

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(decodedBase64); err != nil {
		return err
	}

	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}

func SubmitWaterMeter(file *multipart.FileHeader) ([]int, error) {
	var numbersDetected []int
	src, err := file.Open()

	if err != nil {
		return numbersDetected, err
	}

	defer src.Close()

	fileUploadedExtension := strings.Split(file.Filename, ".")[1]

	timestamp := helpers.GetCurrentTimestampString()
	outputFileName := "room_1_" + timestamp + "." + fileUploadedExtension

	fmt.Println("DEBUG", utilities.GetFilePath(common.WATER_METER_PATH+outputFileName))
	dst, err := os.Create(utilities.GetFilePath(common.WATER_METER_PATH + outputFileName))
	if err != nil {
		return numbersDetected, err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return numbersDetected, err
	}
	texts, err := GetTextDetection(common.WATER_METER_PATH + outputFileName)

	if err != nil {
		return texts, err
	}

	return texts, nil
}
