package helpers

import (
	"encoding/base64"
	"fmt"
	"os"
	"saarm/pkg/common"
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

	timestamp := GetCurrentTimestampString()
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
