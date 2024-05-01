package services

import (
	"encoding/base64"
	"fmt"
	"os"
	"saarm/pkg/common"
	"saarm/pkg/helpers"
	"saarm/pkg/utilities"
	"strings"
)

func storeWaterMeterFile(file common.UploadWaterMeter, roomID string) (string, error) {
	baseData := file.File[strings.IndexByte(file.File, ',')+1:]
	var outputFileName string

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		return outputFileName, err
	}

	parts := strings.SplitN(file.File, ";", 2)
	var fileType string

	if len(parts) != 2 {
		return outputFileName, err
	}

	mimeType := parts[0]
	fileType = strings.Split(mimeType, "/")[1]

	timestamp := helpers.GetCurrentTimestampString()
	outputFileName = "room-" + roomID + "-" + timestamp + "." + fileType

	pathData := utilities.GetFilePath(common.WATER_METER_PATH, outputFileName)

	f, err := os.Create(pathData)

	if err != nil {
		return outputFileName, err
	}

	defer f.Close()

	if _, err := f.Write(decodedBase64); err != nil {
		return outputFileName, err
	}

	if err := f.Sync(); err != nil {
		return outputFileName, err
	}

	fmt.Println("[storeWaterMeterFile] | Created system file")

	return outputFileName, nil
}

func SubmitWaterMeter(file common.UploadWaterMeter, roomID string) ([]string, error) {
	var numbersDetected []string

	fileSubmited, err := storeWaterMeterFile(file, roomID)

	if err != nil {
		return numbersDetected, err
	}

	fmt.Println("SubmitWaterMeter [PATH] : ", common.WATER_METER_PATH, fileSubmited)

	fmt.Println("RUNNING detect water meter")
	numbersDetected, err = GetTextDetection(common.WATER_METER_PATH, fileSubmited)

	if err != nil {
		fmt.Println("RUNNING detect water meter failed: ", err.Error())
		return numbersDetected, err
	}

	return numbersDetected, nil
}
