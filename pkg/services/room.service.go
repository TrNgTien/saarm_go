package services

func SubmitWaterMeter() ([]int, error) {
	texts, err := GetTextDetection()

	if err != nil {
		return texts, err
	}

	return texts, nil
}
