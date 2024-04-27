package helpers

import (
	"strconv"
	"time"
)

func GetOneDay() int64 {
	return time.Now().Add(time.Hour * 24).Unix()
}

func GetCurrentTime() int64 {
	return time.Now().Unix()
}

func GetCurrentTimestampString() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
