package utilities

import "strconv"

func GetIntValue (value string) int {
  v, _ := strconv.Atoi(value)
  return v 
}
