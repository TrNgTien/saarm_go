package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"saarm/pkg/common"
	"saarm/pkg/helpers"
	"saarm/pkg/utilities"
	"strings"
)

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
