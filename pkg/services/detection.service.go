package services

import (
	"context"
	"log"
	"saarm/pkg/utilities"

	vision "cloud.google.com/go/vision/apiv1"
)

func GetTextDetection() ([]int, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	IMAGE_URL := "https://i.imgur.com/wRLnsuU.png"

	image := vision.NewImageFromURI(IMAGE_URL)

	texts, err := client.DetectTexts(ctx, image, nil, 10)

	var rs []int

	for _, text := range texts {
		v, err := utilities.GetIntValue(text.Description)

		if err != nil {

			continue
		}

		rs = append(rs, v)
	}

	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
	}

	return rs, nil
}
