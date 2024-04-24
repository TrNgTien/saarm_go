package services

func SubmitWaterMeter() ([]string, error) {
	texts, err := GetTextDetection()

	if err != nil {
		return texts, err
	}

	return texts, nil
}
