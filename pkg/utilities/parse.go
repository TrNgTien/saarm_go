package utilities

import "strconv"

func GetIntValue(value string) (int, error) {
	v, err := strconv.Atoi(value)
   if err != nil {
     return 0, err
  }

  return v, nil
}
