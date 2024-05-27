package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"saarm/pkg/utilities"

	vision "cloud.google.com/go/vision/apiv1"
)

func GetTextDetection(fileDir, fileName string) ([]string, error) {
	var rs []string
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return rs, err
	}

	defer client.Close()

	IMAGE_WATER_METER_PATH := utilities.GetFilePath(fileDir, fileName)

	fmt.Println("[utilities][GetTextDetection] | open image created!")

	file, err := os.Open(IMAGE_WATER_METER_PATH)

	if err != nil {
		log.Fatalf("[utilities][GetTextDetection] | Failed to read file: %v", err)
		return rs, err
	}

	defer file.Close()

	image, err := vision.NewImageFromReader(file)

	if err != nil {
		log.Fatalf("Failed to detect image: %v", err)
		return rs, err
	}

	texts, err := client.DetectTexts(ctx, image, nil, 10)

	for _, text := range texts {
		_, err := utilities.GetIntValue(text.Description)

		if err != nil {
			continue
		}

		rs = append(rs, text.Description)
	}

	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
		return rs, err
	}

	return rs, nil
}
