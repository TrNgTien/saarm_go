package utilities

import (
	"strconv"

	"github.com/google/uuid"
)

func GetIntValue(value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ParseStringToUuid(value string) uuid.UUID {
	return uuid.MustParse(value)
}
