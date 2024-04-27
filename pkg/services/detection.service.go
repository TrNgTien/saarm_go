package services

import (
	"context"
	"log"
	"os"
	"saarm/pkg/utilities"

	vision "cloud.google.com/go/vision/apiv1"
)

func GetTextDetection(filePath string) ([]int, error) {
	var rs []int
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return rs, err
	}

	defer client.Close()

	IMAGE_WATER_METER_PATH := utilities.GetFilePath(filePath)

	file, err := os.Open(IMAGE_WATER_METER_PATH)

	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
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
		v, err := utilities.GetIntValue(text.Description)

		if err != nil {
			continue
		}

		rs = append(rs, v)
	}

	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
		return rs, err
	}

	return rs, nil
}
