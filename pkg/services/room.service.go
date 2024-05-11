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

func saveFileSystem(file, roomID string) (string, error) {
	baseData := file[strings.IndexByte(file, ',')+1:]
	var outputFileName string

	decodedBase64, err := base64.StdEncoding.DecodeString(baseData)

	if err != nil {
		return outputFileName, err
	}

	parts := strings.SplitN(file, ";", 2)
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

	fileCropped, err := saveFileSystem(file.CroppedFile, roomID)

	if err != nil {
		return numbersDetected, err
	}

	fileOriginal, err := saveFileSystem(file.OriginalFile, roomID)

	if err != nil {
		return numbersDetected, err
	}

	IMAGE_WATER_METER_PATH := utilities.GetFilePath(common.WATER_METER_PATH, fileCropped)
	ORIGINAL_WATER_METER_PATH := utilities.GetFilePath(common.WATER_METER_PATH, fileOriginal)

	info, err := UploadObject(common.MINIO_BUCKET_CROPPED, fileCropped, IMAGE_WATER_METER_PATH)

	if err != nil {
		fmt.Println("Failed to upload images", err.Error())
		return numbersDetected, err
	}

	infoOriginal, err := UploadObject(common.MINIO_BUCKET_ORIGINAL, fileOriginal, ORIGINAL_WATER_METER_PATH)

	if err != nil {
		fmt.Println("Failed to upload images", err.Error())
		return numbersDetected, err
	}

	fmt.Println("Upload image cropped, originalFile success", info, infoOriginal)

	numbersDetected, err = GetTextDetection(common.WATER_METER_PATH, fileCropped)

	if err != nil {
		fmt.Println("RUNNING detect water meter failed: ", err.Error())
		return numbersDetected, err
	}

	return numbersDetected, nil
}

func CreateRoom() error{

  return nil
}

func GetCurrentBill() error{
  return nil
}
