package helpers

import "time"

func GetOneDay() int64 {
	return time.Now().Add(time.Hour * 24).Unix()
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}
