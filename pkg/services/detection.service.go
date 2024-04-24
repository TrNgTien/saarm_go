package services

import (
	"context"
	"log"

	vision "cloud.google.com/go/vision/apiv1"
)

func GetTextDetection() ([]string, error) {
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

	var rs []string

	for _, text := range texts {
		rs = append(rs, text.Description)
	}

	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
	}

	return rs, nil
}
